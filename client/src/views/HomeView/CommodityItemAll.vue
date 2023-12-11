<template>
  <el-table :data="currentTableData" border  show-empty  :key="table_key1">
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
          <el-drawer v-model="drawer" title="商品更多信息" size="50%" destroy-on-close :append-to-body="true" :before-close="handleClose1">
            <div>
              <el-form label-width="120px" style="margin-left: 10px">
                <el-form-item label="生产日期:">
                  <span>{{ this.home_commodity_all.find(item => item.id === this.focus_commodity_item_id).Commodity.produce_at }}</span>
                </el-form-item>
                <el-form-item label="生产地址:">
                  <span>{{ this.home_commodity_all.find(item => item.id === this.focus_commodity_item_id).Commodity.produce_address }}</span>
                </el-form-item>
                <el-form-item label="平台所在国家:">
                  <span>{{ this.home_commodity_all.find(item => item.id === this.focus_commodity_item_id).Platform.country }}</span>
                </el-form-item>
                <el-form-item label="上次更新时间:">
                  <span>{{ this.home_commodity_all.find(item => item.id === this.focus_commodity_item_id).update_at }}</span>
                </el-form-item>
              </el-form>
            </div>
            <div>
              <div style="height: 40px; margin-top: 100px">
                <span style="margin-top: 20px; margin-bottom: 20px">查询价格历史:</span>
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
                  innerDrawer = true;
                  this.find_price_history.commodity_item_id=this.focus_commodity_item_id;
                  findPriceHistory();">查询
                </el-button>
                <el-drawer
                  v-model="innerDrawer"
                  title="价格历史（若数据为空尝试刷新后再查询）"
                  :append-to-body="true"
                  destroy-on-close
                  :before-close="handleClose2"
                >
                  <el-table
                    :data="commodity_price_history"
                    show-empty
                    border
                    style="width: 400px"
                    :row-class-name="highlightLowestPriceRow"
                    :default-sort="{prop: 'update_at', order: 'descending'}"
                  >
                    <el-table-column label="更新时间" prop="update_at" width="200"/>
                    <el-table-column label="价格" prop="new_price" width="200"/>
                  </el-table>
                </el-drawer>
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
    :page-size="this.pageSize"
    layout="total, prev, pager, next, jumper"
    :total="home_commodity_all.length"
    position="bottom"
    background
  />
</template>

<script>

import {Star} from "@element-plus/icons-vue";
import {Plus} from "@element-plus/icons-vue";

export default {
  name: "home_commodity_all",
  components: {Star, Plus},
  data() {
    return {
      home_commodity_all: localStorage.getItem('home_commodity_all')
          ? JSON.parse(localStorage.getItem('home_commodity_all'))
          : [],
      commodity_price_history: localStorage.getItem('commodity_price_history')
        ? JSON.parse(localStorage.getItem('commodity_price_history'))
        : [],
      focus_commodity_item_id:0,
      currentPage: 1,
      pageSize: 10,
      drawer: false,
      innerDrawer: false,
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
      return this.home_commodity_all.slice(start, end);
    },
  },
  created() {
    if(localStorage.getItem('home_commodity_all') === null){
      this.findAll()
    }
  },
  methods: {
    handleCurrentChange(val) {
      this.currentPage = val;
    },
    findAll() {
      this.request.get("/commodity/all").then((res) => {
        if (res.status === 200) {
          localStorage.setItem("home_commodity_all", JSON.stringify(res.data));
        } else {
          this.$message.error(res.message);
        }
      });
    },
    addToFavorite(){
      if(localStorage.getItem("token") === null){
        this.$message.error("请先登录");
      }
      else{
        this.request.post('/favorites',this.focus_commodity_item_id).then(res=>{
          if(res.status === 200){
            this.$message.success("添加成功！")
          }
          else {
            this.$message.error(res.message)
          }
        })
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
            localStorage.setItem("commodity_price_history", JSON.stringify(res.data))
            this.commodity_price_history=localStorage.getItem('commodity_price_history')
            this.showPrice=true
            this.table_key2 = Math.random()
          }
          else {
            this.$message.error(res.message)
          }
        })
        this.innerDrawer = true;
      }
    },
    highlightLowestPriceRow({ row }) {
      const lowestPrice = Math.min(...this.commodity_price_history.map(item => item.new_price));
      return row.new_price === lowestPrice ? 'lowest-price-row' : '';
    },
    handleClose1(){
      this.find_price_history.time_start=null
      this.find_price_history.time_end=null
      this.drawer=false
    },
    handleClose2(){
      this.find_price_history.time_start=null
      this.find_price_history.time_end=null
      this.innerDrawer=false
      this.commodity_price_history=[]
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