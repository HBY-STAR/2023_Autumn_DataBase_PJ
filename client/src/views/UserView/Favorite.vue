<template>
  <el-table :data="user_favorite" border height="580px" show-empty>
    <el-table-column label="商品序号" prop="commodity_item_id" width="160"/>
    <el-table-column label="商品名" prop="commodityItem.item_name" width="160"/>
    <el-table-column label="商家" prop="commodityItem.seller.username" width="150"/>
    <el-table-column label="平台" prop="commodityItem.platform.name" width="150"/>
    <el-table-column label="当前价格" prop="commodityItem.price" width="150"/>
    <el-table-column label="更多信息" width="100">
      <template v-slot="scope">
        <div style="text-align: center">
          <el-icon>
            <Search style="height: 25px; width: 25px;text-align: center;color: #7300ff" @click="this.focus_commodity_item_id=scope.row.commodity_item_id; drawer = true;">
            </Search>
          </el-icon>
        </div>
        <el-drawer v-model="drawer" title="商品更多信息" size="50%">
          <div>
            <el-form label-width="80px" size="small" style="margin-left: 10px">
              <el-form-item label="生产日期:">
                <span>{{ user_favorite.commodity.produce_at }}</span>
              </el-form-item>
              <el-form-item label="生产地址:">
                <span>{{ user_favorite.commodity.produce_address }}</span>
              </el-form-item>
              <el-form-item label="售卖平台所属国家:">
                <span>{{ user_favorite.platform.country }}</span>
              </el-form-item>
              <el-form-item label="上次信息更新时间:">
                <span>{{ user_favorite.update }}</span>
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
    <el-table-column label="收藏时间" prop="create_at" width="150"/>
    <el-table-column label="提醒价格" prop="price_limit" width="120"/>
    <el-table-column label="设置提醒价格" width="110">
      <template v-slot="scope">
        <div style="text-align: center">
          <el-icon>
            <Edit style="height: 25px; width: 25px;text-align: center;color: #409eff" @click="this.focus_commodity_item_id=scope.row.commodity_item_id; dialogFormVisible=true;">
            </Edit>
          </el-icon>
        </div>
        <el-dialog v-model="dialogFormVisible" title="设置提醒价格">
          <el-input-number v-model="input_limit" :min="0" :precision="2"/>
          <template #footer>
            <span class="dialog-footer">
              <el-button @click="dialogFormVisible = false">返回</el-button>
              <el-button type="primary" @click="dialogFormVisible = false; setPriceLimit();">
                确定
              </el-button>
            </span>
          </template>
        </el-dialog>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>

import { Edit, Search } from "@element-plus/icons-vue";

export default {
  name: "user_favorite",
  components: { Search, Edit },
  data() {
    return {
      user_favorite: localStorage.getItem('user_favorite')
        ? JSON.parse(localStorage.getItem('user_favorite'))
        : [],
      commodity_price_history: localStorage.getItem('commodity_price_history')
        ? JSON.parse(localStorage.getItem('commodity_price_history'))
        : [],

      focus_commodity_item_id: 0,
      set_price_limit:{
        item_id: -1,
        price_limit: -1,
      },
      dialogFormVisible:false,
      input_limit: -1,

      drawer: false,
      innerDrawer: false,
      find_price_history:{
        commodity_item_id: -1,
        time_start: null,
        time_end: null,
      },
    };
  },
  computed: {
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
    this.findAll()
  },
  methods: {
    findAll(){
      this.request.get("/favorites/all").then((res) => {
        if (res.code === "200") {
          localStorage.setItem("user_favorite", JSON.stringify(res.data));
        } else {
          this.$message.error(res.message);
        }
      });
    },
    setPriceLimit(){
      if(this.input_limit<0){
        this.$message.warning("提醒价格的值不应低于0")
        return
      }
      this.set_price_limit.item_id=this.focus_commodity_item_id
      this.set_price_limit.price_limit=this.input_limit
      this.request.post("/price/limit",this.set_price_limit).then((res) => {
        if (res.code === "200") {
          this.$message.success("设置成功")
        } else {
          this.$message.error(res.message);
        }
      });
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
    },
  }
};

</script>

<style>
.highlight {
  background-color: #ff0036; /* 设置您想要的高亮颜色 */
}
</style>