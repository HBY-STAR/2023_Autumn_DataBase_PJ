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
          <el-option label="精确查找" value= true />
          <el-option label="模糊查找" value = false />
        </el-select>
      </template>
      <template #append>
        <el-button @click="search_commodity">
          <el-icon><Search /></el-icon>
        </el-button>
      </template>
    </el-input>
  </div>
  <el-table :data="currentTableData" border style="margin-top: 20px" show-empty>
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
              <div style="height: 40px; margin-top: 100px">
                <span style="margin-top: 20px; margin-bottom: 20px">查询价格历史:</span>
              </div>
              <div style="height: 50px">
                <el-date-picker
                  v-model="find_price_history.time_start"
                  type="date"
                  value-format="YYYY-MM-DD hh:mm:ss"
                  placeholder="起始时间"
                />
              </div>
              <div style="height: 50px">
                <el-date-picker
                  v-model="find_price_history.time_end"
                  type="date"
                  value-format="YYYY-MM-DD hh:mm:ss"
                  placeholder="结束时间"
                />
              </div>
              <div>
                <el-button @click="
                  this.find_price_history.commodity_item_id=this.focus_commodity_item_id;
                  findPriceHistory();
                  innerDrawer = true;
                  ">查询
                </el-button>
                <el-drawer
                  v-model="innerDrawer"
                  title="价格历史"
                  :append-to-body="true"
                  destroy-on-close
                  :before-close="handleClose2"
                >
                  <el-table
                    :data="commodity_price_history"
                    border
                    show-empty
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
      search_commodity_item: localStorage.getItem('search_commodity_item')
        ? JSON.parse(localStorage.getItem('search_commodity_item'))
        : [],
      commodity_price_history: localStorage.getItem('commodity_price_history')
        ? JSON.parse(localStorage.getItem('commodity_price_history'))
        : [],
      focus_commodity_item_id:0,
      currentPage: 1,
      pageSize: 10,
      search_info: {
        search: null,
        accurate:  null,
        range: null,
      },
      drawer: false,
      innerDrawer: false,
      showPrice:false,
      find_price_history:{
        commodity_item_id: -1,
        time_start: null,
        time_end: null,
      },
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
        })
      }
    },
    search_commodity(){
      if(this.search_info.search==null || this.search_info.accurate==null || this.search_info.range == null){
        this.$message.error('请确定范围选项并输入要搜索的内容');
      }else {
        this.search_info.accurate=Boolean(this.search_info.accurate)
        this.request.post("/search",this.search_info).then((res) => {
          if (res.status === 200) {
            localStorage.setItem("search_commodity_item", JSON.stringify(res.data));
            this.search_commodity_item = JSON.stringify(res.data)
            location.reload();
          } else {
            this.$message.error(res.message);
          }
        });
      }
    },
    findPriceHistory(){
      if(this.find_price_history.commodity_item_id===-1){
        this.$message.error('未选中任何商品')
      }else {
        this.request.post('/price/history',this.find_price_history).then(res=>{
          if(res.status===200){
            localStorage.setItem("commodity_price_history", JSON.stringify(res.data))
            this.commodity_price_history = JSON.stringify(res.data)
            this.showPrice=true
          }
          else {
            this.$message.error(res.message)
          }
        })
      }
    },
    highlightLowestPriceRow({ row }) {
      const lowestPrice = Math.min(...this.commodity_price_history.map(item => item.new_price));
      return row.new_price === lowestPrice ? 'lowest-price-row' : '';
    },
    handleClose1(){
      this.drawer=false
    },
    handleClose2(){
      this.innerDrawer=false
      this.showPrice=false
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