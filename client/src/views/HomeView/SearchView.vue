<template>
  <div>
    <el-input
      v-model="search_info.search"
      placeholder="请输入..."
    >
      <template #prepend>
        <el-select v-model="search_info.accurate" placeholder="搜索" style="width: 120px">
          <el-option label="按相关商品" value = false />
          <el-option label="按具体种类" value= true />
        </el-select>
      </template>
      <template #append>
        <el-button @click="search_commodity">
          <el-icon><Search /></el-icon>
        </el-button>
      </template>
    </el-input>
  </div>
  <el-table :data="currentTableData" border height="480px" show-empty style="margin-top: 20px">
    <el-table-column label="商品序号" prop="commodity_id" width="180"/>
    <el-table-column label="商品名" prop="item_name" width="200"/>
    <el-table-column label="种类" prop="commodity.category" width="200"/>
    <el-table-column label="价格" prop="price" width="200"/>
    <el-table-column label="商家" prop="seller.username" width="150"/>
    <el-table-column label="平台" prop="platform.name" width="150"/>
    <el-table-column label="更多信息" width="100">
      <template v-slot="scope">
        <div style="text-align: center">
          <el-icon>
            <Search style="height: 25px; width: 25px;text-align: center;color: #7300ff" @click="this.focus_commodity_item_id=scope.row.commodity_id; drawer = true;">
            </Search>
          </el-icon>
        </div>
        <el-drawer v-model="drawer" title="商品更多信息" size="50%">
          <div>
            <el-form label-width="80px" size="small" style="margin-left: 10px">
              <el-form-item label="生产日期:">
                <span>{{ search_commodity_item.commodity.produce_at }}</span>
              </el-form-item>
              <el-form-item label="生产地址:">
                <span>{{ search_commodity_item.commodity.produce_address }}</span>
              </el-form-item>
              <el-form-item label="售卖平台所属国家:">
                <span>{{ search_commodity_item.platform.country }}</span>
              </el-form-item>
              <el-form-item label="上次信息更新时间:">
                <span>{{ search_commodity_item.update }}</span>
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
            <Star style="height: 25px; width: 25px;text-align: center;color: #409eff" @click="this.focus_commodity_item_id=scope.row.commodity_id; addToFavorite();">
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

import { Search, Star } from "@element-plus/icons-vue";

export default {
  name: "search_commodity_item",
  components: {Star, Search},
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
    search_commodity(){
      if(this.search_info.search==null || this.search_info.accurate==null){
        this.$message.error('请确定搜索选项并输入要搜索的内容');
      }else {
        this.request.post("/search",).then((res) => {
          if (res.code === "200") {
            localStorage.setItem("search_commodity_item", JSON.stringify(res.data));
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
          if(res.code==='200'){
            this.$message.success("添加成功！")
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
    }
  }
}
</script>

<style>
.highlight {
  background-color: #ff0036; /* 设置您想要的高亮颜色 */
}
</style>