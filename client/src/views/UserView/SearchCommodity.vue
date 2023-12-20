<template>
  <div>
    <el-input
      v-model="search_info.search"
      placeholder="请输入..."
    >
      <template #prepend>
        <el-select v-model="search_info.range" placeholder="范围" style="width: 100px;">
          <el-option label="按名称" value = 'name' />
          <el-option label="按种类" value= 'category' />
        </el-select>
        <el-select v-model="search_info.accurate" placeholder="选项"  style="width: 100px;margin-left: 20px">
          <el-option label="精确查找" value= 'true' />
          <el-option label="模糊查找" value = 'false' />
        </el-select>
      </template>
      <template #append>
        <el-button @click="SearchCommodity">
          <el-icon><Search /></el-icon>
        </el-button>
      </template>
    </el-input>
  </div>
  <el-table :data="currentTableData" border show-empty :key="table_key1" style="width: 1000px; margin-top: 20px">
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
                {{ this.search_commodity.find(item => item.id === this.focus_commodity_id).default_name }}
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
    :total="search_commodity.length"
    position="bottom"
    background
  />
</template>

<script>

import { Plus, Search} from "@element-plus/icons-vue";

export default {
  name: "search_commodity_item",
  components: {Plus, Search},
  data() {
    return {
      search_commodity: [],
      find_item:[],
      focus_commodity_id:0,
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
      //date
      range: [],
    }
  },
  computed: {
    currentTableData() {
      const start = (this.currentPage - 1) * this.pageSize;
      const end = start + this.pageSize;
      return this.search_commodity.slice(start, end);
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
  },
  methods: {
    handleCurrentChange(val) {
      this.currentPage = val;
    },
    SearchCommodity(){
      if(this.search_info.search==null || this.search_info.accurate==null || this.search_info.range == null){
        this.$message.error('请确定范围选项并输入要搜索的内容');
      }else {
        this.search_info.accurate = this.search_info.accurate === 'true'
        this.request.post("/commodities/search",this.search_info).then((res) => {
          if (res.status === 200) {
            this.search_commodity= res.data
          }
        }).catch(error => {
          this.$message.error(error.response.data.message);
        });
      }
    },
    highlightLowestPriceRow({ row }) {
      const lowestPrice = Math.min(...this.sortedItems.map(item => item.price));
      return row.price === lowestPrice ? "lowest-price-row" : "";
    },
    handleClose1(){
      this.find_item=[];
      this.drawer=false;
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
}
</script>

<style>
.el-table .lowest-price-row {
  background-color: rgba(220, 29, 29, 0.74); /* 设置最低价格行的背景色 */
  /* 可以根据需要添加其他样式 */
}

</style>