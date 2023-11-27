<template>
  <el-table :data="currentTableData" border height="500px" show-empty>
    <el-table-column label="商品序号" prop="commodity_id" width="100"/>
    <el-table-column label="商品名" prop="item_name" width="100"/>
    <el-table-column label="种类" prop="commodity.category" width="100"/>
    <el-table-column label="价格" prop="price" width="100"/>
    <el-table-column label="商家" prop="seller.username" width="100"/>
    <el-table-column label="平台" prop="platform.name" width="100"/>
    <el-table-column label="收藏" width="120">
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
    :total="home_commodity_all.length"
    position="bottom"
    background
  />
</template>

<script>

import {Star} from "@element-plus/icons-vue";

export default {
  name: "home_commodity_all",
  components: {Star},
  data() {
    return {
      home_commodity_all: localStorage.getItem('home_commodity_all')
          ? JSON.parse(localStorage.getItem('home_commodity_all'))
          : [],

      focus_commodity_item_id:0,
      currentPage: 1,
      pageSize: 10,
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
      //this.findAll()
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
      if(localStorage.getItem("token")==null){
        this.$message.error("请先登录");
      }else{
        this.request.post('favorites',this.focus_commodity_item_id).then(res=>{
          if(res.code==='200'){
            this.$message.success("添加成功！")
          }
          else {
            this.$message.error(res.message)
          }
        })
      }

    }
  }
}
</script>