<template>
  <el-table :data="currentTableData" border show-empty :key="table_key1" style="width: 1000px">
    <el-table-column label="商品ID" prop="id" width="120"/>
    <el-table-column label="默认名称" prop="default_name" width="150"/>
    <el-table-column label="种类" prop="category" width="150"/>
    <el-table-column label="生产日期" prop="produce_at" width="200"/>
    <el-table-column label="生产地址" prop="produce_address" width="280"/>
    <el-table-column label="商品实例" width="100">
      <template v-slot="scope">
        <div style="text-align: center">
          <el-icon>
            <Plus style="height: 25px; width: 25px;text-align: center;color: #7300ff"
                  @click="this.focus_commodity_id=scope.row.id; drawer = true; findItem();">
            </Plus>
          </el-icon>
          <el-drawer v-model="drawer" title="商品实例" size="50%" destroy-on-close :before-close="handleClose1"
                     :append-to-body="true">
            <div>
              <span style="margin-left: 20px; color: #007dff">
                {{ this.user_commodity_all.find(item => item.id === this.focus_commodity_id).default_name }}
              </span>
            </div>
            <div>
              <el-table :data="sortedItems" border height="500px" show-empty style="margin-top: 30px" :row-class-name="highlightLowestPriceRow">
                <el-table-column label="商品序号" prop="id" width="120" sortable="custom"/>
                <el-table-column label="商品名" prop="item_name" width="150" />
                <el-table-column label="种类" prop="Commodity.category" width="100" />
                <el-table-column label="价格" prop="price" width="100" />
                <el-table-column label="商家" prop="Seller.username" width="150" />
                <el-table-column label="平台" prop="Platform.name" width="107" />
              </el-table>
            </div>
          </el-drawer>
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
    :total="user_commodity_all.length"
    position="bottom"
    background
  />
</template>

<script>

import { Plus} from "@element-plus/icons-vue";

export default {
  name: "user_commodity_all",
  components: { Plus },
  data() {
    return {
      user_commodity_all: localStorage.getItem("user_commodity_all")
        ? JSON.parse(localStorage.getItem("user_commodity_all"))
        : [],
      focus_commodity_id: 0,
      currentPage: 1,
      pageSize: 10,
      drawer: false,
      //key
      table_key1: "",
      find_item:[],
    };
  },
  computed: {
    currentTableData() {
      const start = (this.currentPage - 1) * this.pageSize;
      const end = start + this.pageSize;
      return this.user_commodity_all.slice(start, end);
    },
    sortedItems() {
      if(this.find_item.length>0){
        return this.find_item.slice().sort((a, b) => a.id - b.id);
      }else {
        return []
      }
    },
  },
  created() {
    if (localStorage.getItem("user_commodity_all") === null) {
      this.findAll();
    }
  },
  methods: {
    handleCurrentChange(val) {
      this.currentPage = val;
    },
    findAll() {
      this.request.get("/commodities/all").then((res) => {
        if (res.status === 200) {
          localStorage.setItem("user_commodity_all", JSON.stringify(res.data));
        } else {
          this.$message.error(res.message);
        }
      }).catch(error => {
        this.$message.error(error.response.data.message);
      });
    },
    highlightLowestPriceRow({ row }) {
      const lowestPrice = Math.min(...this.sortedItems.map(item => item.price));
      return row.price === lowestPrice ? "lowest-price-row" : "";
    },
    handleClose1() {
      this.find_item = [];
      this.drawer = false;
    },
    findItem(){
      this.request.get("/item/commodity/"+this.focus_commodity_id).then((res) => {
        if (res.status === 200) {
          this.find_item=res.data
        }
      }).catch(error => {
        this.$message.error(error.response.data.message);
      });
    }
  }
};
</script>

<style>
.el-table .lowest-price-row {
  background-color: rgba(220, 29, 29, 0.74); /* 设置最低价格行的背景色 */
  /* 可以根据需要添加其他样式 */
}

</style>