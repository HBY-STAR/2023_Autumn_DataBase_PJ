<template v-slot="scope">
  <div style="text-align: center">
    <el-button>
      <el-icon>
        <Search style="height: 25px; width: 25px;text-align: center;color: aqua" @click="drawer = true;">
        </Search>
      </el-icon>
    </el-button>
  </div>
  <el-drawer v-model="drawer" title="商品更多信息" size="50%">
    <div>
<!--      <el-form label-width="80px" size="small" style="margin-left: 10px">-->
<!--        <el-form-item label="生产日期:">-->
<!--          <span>{{ search_commodity_item.commodity.produce_at }}</span>-->
<!--        </el-form-item>-->
<!--        <el-form-item label="生产地址:">-->
<!--          <span>{{ search_commodity_item.commodity.produce_address }}</span>-->
<!--        </el-form-item>-->
<!--        <el-form-item label="售卖平台所属国家:">-->
<!--          <span>{{ search_commodity_item.platform.country }}</span>-->
<!--        </el-form-item>-->
<!--        <el-form-item label="上次信息更新时间:">-->
<!--          <span>{{ search_commodity_item.update }}</span>-->
<!--        </el-form-item>-->
<!--      </el-form>-->
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
          <el-table :data="commodity_price_history" border show-empty style="width: 400px">
            <el-table-column label="更新时间" prop="update_at" width="200"/>
            <el-table-column label="价格" prop="new_price" width="200"/>
          </el-table>
        </el-drawer>
      </div>
    </div>
  </el-drawer>
</template>

<script>

// eslint-disable-next-line no-unused-vars
import { Search } from "@element-plus/icons-vue";

export default {
  name: "search_commodity_item",
  components: {Search},
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
      }

    }
  },
  computed: {
    Search() {
      return Search
    },
    currentTableData() {
      const start = (this.currentPage - 1) * this.pageSize;
      const end = start + this.pageSize;
      return this.search_commodity_item.slice(start, end);
    },
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
    }
  }
}
</script>


