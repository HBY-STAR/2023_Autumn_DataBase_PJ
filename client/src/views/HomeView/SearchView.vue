<template>
  <div>
    <el-input
      v-model="search_info.search"
      placeholder="请输入..."
    >
      <template #prepend>
        <el-select v-model="search_info.range" placeholder="范围" style="width: 100px;">
          <el-option label="按商品" value = 'name' />
          <el-option label="按种类" value= 'category' />
        </el-select>
        <el-select v-model="search_info.accurate" placeholder="选项"  style="width: 100px;margin-left: 20px">
          <el-option label="精确查找" value= 'true' />
          <el-option label="模糊查找" value = 'false' />
        </el-select>
      </template>
      <template #append>
        <el-button @click="search_commodity">
          <el-icon><Search /></el-icon>
        </el-button>
      </template>
    </el-input>
  </div>
  <el-table :data="currentTableData" border style="margin-top: 20px" show-empty  :key="table_key1">
    <el-table-column label="商品序号" prop="id" width="180"/>
    <el-table-column label="商品名" prop="item_name" width="200"/>
    <el-table-column label="种类" prop="Commodity.category" width="200"/>
    <el-table-column label="价格" prop="price" width="200"/>
    <el-table-column label="商家" prop="Seller.username" width="150"/>
    <el-table-column label="平台" prop="Platform.name" width="150"/>
    <el-table-column label="更多信息" width="100">
      <template v-slot="scope">
        <div style="text-align: center">
          <el-icon>
            <Plus style="height: 25px; width: 25px;text-align: center;color: #7300ff" @click="this.focus_commodity_item_id=scope.row.id; drawer = true">
            </Plus>
          </el-icon>
          <el-drawer v-model="drawer" title="更多信息" size="30%" destroy-on-close :append-to-body="true" :before-close="handleClose1">
            <div>
              <el-form label-width="120px" style="margin-left: 10px">
                <el-form-item label="生产日期:">
                  <span>{{ this.search_commodity_item.find(item => item.id === this.focus_commodity_item_id).Commodity.produce_at }}</span>
                </el-form-item>
                <el-form-item label="生产地址:">
                  <span>{{ this.search_commodity_item.find(item => item.id === this.focus_commodity_item_id).Commodity.produce_address }}</span>
                </el-form-item>
                <el-form-item label="平台所在国家:">
                  <span>{{this.search_commodity_item.find(item => item.id === this.focus_commodity_item_id).Platform.country }}</span>
                </el-form-item>
                <el-form-item label="上次更新时间:">
                  <span>{{ this.search_commodity_item.find(item => item.id === this.focus_commodity_item_id).update_at }}</span>
                </el-form-item>
              </el-form>
            </div>
            <div>
              <div style="height: 40px; margin-top: 20px">
                <span style="margin-top: 20px; margin-bottom: 20px; color: #007dff">查询价格历史:</span>
              </div>
              <el-date-picker style="margin-bottom: 20px"
                              v-model="range"
                              type="daterange"
                              unlink-panels
                              range-separator="到"
                              start-placeholder="起始时间"
                              end-placeholder="结束时间"
                              value-format="YYYY-MM-DD hh:mm:ss"
                              :shortcuts="shortcuts"
              />
              <div>
                <el-button @click="
                  this.find_price_history.commodity_item_id=this.focus_commodity_item_id;
                  findPriceHistory();
                ">查询
                </el-button>
                <el-table
                  :data="commodity_price_history"
                  show-empty
                  border
                  style="width: 400px;margin-top: 20px; height: 250px"
                  :row-class-name="highlightLowestPriceRow"
                  :default-sort="{prop: 'update_at', order: 'descending'}"
                  :key="table_key2"
                >
                  <el-table-column label="更新时间" prop="update_at" width="200"/>
                  <el-table-column label="价格" prop="new_price" width="200"/>
                </el-table>

                <div style="width: 400px;height: 300px">
                  <canvas ref="lineChart" width="400" height="300" style="margin-top: 20px;"></canvas>
                </div>
              </div>
            </div>
          </el-drawer>
        </div>
      </template>
    </el-table-column>
    <el-table-column label="收藏" width="100">
      <template v-slot="scope">
        <div style="text-align: center">
          <el-icon>
            <Star style="height: 25px; width: 25px;text-align: center;color: #409eff" @click="this.focus_commodity_item_id=scope.row.id; addToFavorite();">
            </Star>
          </el-icon>
        </div>
      </template>
    </el-table-column>
  </el-table>
  <el-pagination
    style="margin-top: 20px"
    @current-change="handleCurrentChange"
    :current-page="currentPage"
    :page-size="20"
    layout="total, prev, pager, next, jumper"
    :total="search_commodity_item.length"
    position="bottom"
    background
  />
</template>

<script>

import { Plus, Search, Star } from "@element-plus/icons-vue";

export default {
  name: "search_commodity_item",
  components: {Star, Plus, Search},
  data() {
    return {
      search_commodity_item: [],
      commodity_price_history:[],
      focus_commodity_item_id:0,
      currentPage: 1,
      pageSize: 10,
      search_info: {
        search: null,
        accurate:  null,
        range: null,
      },
      drawer: false,
      showPrice:false,
      //key
      table_key1:'',
      table_key2:'',
      find_price_history:{
        commodity_item_id: -1,
        time_start: null,
        time_end: null,
      },
      //date
      range: [],
      shortcuts:[
        {
          text: '近一周',
          value: function () {
            const end = new Date();
            const start = new Date();
            start.setDate(end.getDate() - 7);
            return [start, end];
          },
        },
        {
          text: '近一月',
          value: function () {
            const end = new Date();
            const start = new Date();
            start.setMonth(end.getMonth() - 1);
            return [start, end];
          },
        },
        {
          text: '近一年',
          value: function () {
            const end = new Date();
            const start = new Date();
            start.setFullYear(end.getFullYear() - 1);
            return [start, end];
          },
        },
      ],
    }
  },
  computed: {
    currentTableData() {
      const start = (this.currentPage - 1) * this.pageSize;
      const end = start + this.pageSize;
      return this.search_commodity_item.slice(start, end);
    },
    lowestPriceRowIndex() {
      let lowestPrice = Number.MAX_VALUE;
      let lowestIndex = -1;

      this.commodity_price_history.forEach((item, index) => {
        if (item.new_price < lowestPrice) {
          lowestPrice = item.new_price;
          lowestIndex = index;
        }
      });

      return lowestIndex;
    }
  },
  created() {
  },
  watch: {
    commodity_price_history() {
      // 监听数据变化，重新绘制图表
      this.drawChart();
    }
  },
  mounted() {
    // 组件挂载后立即绘制图表
    this.drawChart();
  },
  methods: {
    handleCurrentChange(val) {
      this.currentPage = val;
    },
    addToFavorite(){
      if(localStorage.getItem("token")===null){
        this.$message.error("请先登录");
      }
      else{
        this.request.post('/favorites',this.focus_commodity_item_id).then(res=>{
          if(res.status===200){
            this.$message.success("添加成功！")
          }
          else {
            this.$message.error(res.message)
          }
        }).catch(error => {
          this.$message.error(error.response.data.message);
        });
      }
    },
    search_commodity(){
      if(this.search_info.search==null || this.search_info.accurate==null || this.search_info.range == null){
        this.$message.error('请确定范围选项并输入要搜索的内容');
      }else {
        this.search_info.accurate = this.search_info.accurate === 'true'
        this.request.post("/search",this.search_info).then((res) => {
          if (res.status === 200) {
            this.search_commodity_item= res.data
          } else {
            this.$message.error(res.message);
          }
        }).catch(error => {
          this.$message.error(error.response.data.message);
        });
      }
    },
    findPriceHistory(){
      this.find_price_history.commodity_item_id=this.focus_commodity_item_id
      this.find_price_history.time_start=this.range[0]
      this.find_price_history.time_end=this.range[1]
      if(this.find_price_history.commodity_item_id===-1){
        this.$message.error('未选中任何商品')
      }else {
        this.request.post('/price/history',this.find_price_history).then(res=>{
          if(res.status===200){
            this.commodity_price_history = res.data
            this.table_key2 = Math.random()
          }
          else {
            this.$message.error(res.message)
          }
        }).catch(error => {
          this.$message.error(error.response.data.message);
        });
      }
    },
    highlightLowestPriceRow({ row }) {
      const lowestPrice = Math.min(...this.commodity_price_history.map(item => item.new_price));
      return row.new_price === lowestPrice ? 'lowest-price-row' : '';
    },
    handleClose1(){
      this.find_price_history.time_start=null
      this.find_price_history.time_end=null
      this.commodity_price_history=[]
      this.drawer=false
    },
    drawChart() {
      const canvas = this.$refs.lineChart;
      if (!canvas) {
        // 确保 canvas 存在
        return;
      }
      const ctx = canvas.getContext("2d");
      // 清空画布
      ctx.clearRect(0, 0, canvas.width, canvas.height);

      const data = this.commodity_price_history;

      if (data.length === 0) {
        // 如果数据为空，显示提示信息或者不进行绘制
        // 可以添加代码显示“暂无数据”之类的提示
        return;
      }
      // 设置字体样式和位置
      ctx.font = "12px Arial";
      ctx.fillStyle = "black";
      // 绘制 x 轴标题
      ctx.fillText("时间", canvas.width / 2, canvas.height - 10);
      // 绘制 y 轴标题，需要旋转文字
      ctx.save();
      ctx.translate(20, canvas.height / 2);
      ctx.rotate(-Math.PI / 2);
      ctx.textAlign = "right";
      ctx.fillText("价格", 0, 0);
      ctx.restore();
      // 获取更新时间和价格数组
      const updateTimes = data.map(item => item.update_at);
      const prices = data.map(item => item.new_price);
      // 绘制坐标轴
      const xAxis = 30; // X轴起点横坐标
      const yAxis = 270; // Y轴起点纵坐标
      const yMax = Math.max(...prices);
      const yMin = Math.min(...prices);
      const yRange = yMax - yMin;
      const xInterval = (canvas.width - xAxis) / updateTimes.length;
      const yInterval = (yAxis - 50) / yRange;
      // 绘制X轴和Y轴
      ctx.beginPath();
      ctx.strokeStyle = 'black';
      // 绘制X轴线段
      ctx.moveTo(xAxis, yAxis);
      ctx.lineTo(canvas.width - 20, yAxis);
      // 绘制X轴箭头
      ctx.lineTo(canvas.width - 25, yAxis - 5); // 右上角
      ctx.moveTo(canvas.width - 20, yAxis);
      ctx.lineTo(canvas.width - 25, yAxis + 5); // 右下角
      // 绘制Y轴线段
      ctx.moveTo(xAxis, yAxis);
      ctx.lineTo(xAxis, 50);
      // 绘制Y轴箭头
      ctx.lineTo(xAxis - 5, 55); // 左下角
      ctx.moveTo(xAxis, 50);
      ctx.lineTo(xAxis + 5, 55); // 右下角
      // 执行绘制
      ctx.stroke();
      // 绘制价格折线
      ctx.beginPath();
      ctx.strokeStyle = "green";
      ctx.lineWidth = 2;
      updateTimes.forEach((time, index) => {
        const x = xAxis + xInterval * index;
        const y = yAxis - ((prices[index] - yMin) * yInterval);
        if (index === 0) {
          ctx.moveTo(x, y);
        } else {
          ctx.lineTo(x, y);
        }
        ctx.fillStyle = "blue";
        ctx.fillText(this.commodity_price_history.find(item => item.update_at === time).new_price, x, y - 10);
        // 绘制更新时间标签
        //ctx.fillText(time, x - 10, yAxis + 10);
      });
      ctx.stroke();
    },
  }
}
</script>

<style>
.el-table .lowest-price-row {
  background-color: #fa0404; /* 设置最低价格行的背景色 */
  /* 可以根据需要添加其他样式 */
}
</style>