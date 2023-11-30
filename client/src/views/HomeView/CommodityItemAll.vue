<template>
  <el-table :data="currentTableData" border height="530px" show-empty>
    <el-table-column label="商品序号" prop="id" width="180"/>
    <el-table-column label="商品名" prop="item_name" width="200"/>
    <el-table-column label="种类" prop="commodity.category" width="200"/>
    <el-table-column label="价格" prop="price" width="200"/>
    <el-table-column label="商家" prop="seller.username" width="150"/>
    <el-table-column label="平台" prop="platform.name" width="150"/>
    <el-table-column label="更多信息" width="100">
      <template v-slot="scope">
        <div style="text-align: center">
          <el-icon>
            <Search style="height: 25px; width: 25px;text-align: center;color: #7300ff" @click="this.focus_commodity_item_id=scope.row.id; drawer = true;">
            </Search>
          </el-icon>
        </div>
        <el-drawer v-model="drawer" title="商品更多信息" size="50%">
          <div>
            <el-form label-width="80px" size="small" style="margin-left: 10px">
              <el-form-item label="生产日期:">
                <span>{{ home_commodity_all.commodity.produce_at }}</span>
              </el-form-item>
              <el-form-item label="生产地址:">
                <span>{{ home_commodity_all.commodity.produce_address }}</span>
              </el-form-item>
              <el-form-item label="售卖平台所属国家:">
                <span>{{ home_commodity_all.platform.country }}</span>
              </el-form-item>
              <el-form-item label="上次信息更新时间:">
                <span>{{ home_commodity_all.update }}</span>
              </el-form-item>
            </el-form>
          </div>
          <div>
            <div style="height: 40px; margin-top: 20px">
              <span style="margin-top: 20px; margin-bottom: 20px">查询价格历史:</span>
            </div>
            <div style="height: 50px">
              <el-date-picker
                v-model="find_price_history.time_start"
                type="date"
                placeholder="起始时间"
              />
              <el-date-picker
                v-model="find_price_history.time_end"
                type="date"
                placeholder="结束时间"
              />
            </div>
            <div>
              <el-button @click="
            innerDrawer = true;
            this.find_price_history.commodity_item_id=this.focus_commodity_item_id;
            findPriceHistory();">查询
              </el-button>
              <el-drawer
                v-model="innerDrawer"
                title="价格历史"
                :append-to-body="true"
              >
                <p>_(:зゝ∠)_</p>
                <el-table :data="commodity_price_history" border show-empty style="width: 400px" :row-class-name="highlightLowestPriceRow">
                  <el-table-column label="更新时间" prop="update_at" width="200"/>
                  <el-table-column label="价格" prop="new_price" width="200"/>
                </el-table>
              </el-drawer>
            </div>
          </div>
        </el-drawer>
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
    :total="home_commodity_all.length"
    position="bottom"
    background
  />
</template>

<script>

import {Star} from "@element-plus/icons-vue";
import {Search} from "@element-plus/icons-vue";



export default {
  name: "home_commodity_all",
  components: {Star, Search},
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
      return this.home_commodity_all.slice(start, end);
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
        if (res.code === "200") {
          localStorage.setItem("home_commodity_all", JSON.stringify(res.data));
        } else {
          this.$message.error(res.message);
        }
      });
    },
    addToFavorite(){
      if(localStorage.getItem("token")===null){
        this.$message.error("请先登录");
      }else if(localStorage.getItem("user_type") !== 'user')
      {
        this.$message.error("仅普通用户可收藏");
      }
      else{
        this.request.post('/favorites',this.focus_commodity_item_id).then(res=>{
          if(res.code==='200'){
            this.$message.success("添加成功！")
          }
          else {
            this.$message.error(res.message)
          }
        })
      }
    },
    findPriceHistory(){
      if(this.find_price_history.commodity_item_id===-1){
        this.$message.error('未选中任何商品')
      }else {
        this.request.post('/price/history',this.find_price_history).then(res=>{
          if(res.code==='200'){
            localStorage.setItem("commodity_price_history", JSON.stringify(res.data));
          }
          else {
            this.$message.error(res.message)
          }
        })
      }
    },
    highlightLowestPriceRow(row, index) {
      if (index === this.lowestPriceRowIndex) {
        return 'highlight'; // Apply the 'highlight' class to the row with the lowest price
      }
      return ''; // Return empty string for other rows
    },
  }
}
</script>

<style>
.highlight {
  background-color: #ff0036; /* 设置您想要的高亮颜色 */
}
</style>