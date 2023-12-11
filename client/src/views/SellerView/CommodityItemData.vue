<template>
  <el-table :data="currentTableData" border show-empty :key="table_key1">
    <el-table-column label="商品序号" prop="id" width="180"/>
    <el-table-column label="商品名" prop="item_name" width="200"/>
    <el-table-column label="种类" prop="Commodity.category" width="200"/>
    <el-table-column label="价格" prop="price" width="200"/>
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
                  <span>{{ this.seller_commodity_item.find(item => item.id === this.focus_commodity_item_id).Commodity.produce_at }}</span>
                </el-form-item>
                <el-form-item label="生产地址:">
                  <span>{{ this.seller_commodity_item.find(item => item.id === this.focus_commodity_item_id).Commodity.produce_address }}</span>
                </el-form-item>
                <el-form-item label="平台所在国家:">
                  <span>{{ this.seller_commodity_item.find(item => item.id === this.focus_commodity_item_id).Platform.country }}</span>
                </el-form-item>
                <el-form-item label="上次更新时间:">
                  <span>{{ this.seller_commodity_item.find(item => item.id === this.focus_commodity_item_id).update }}</span>
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
                <el-button @click="findPriceHistory">查询</el-button>
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
    <el-table-column label="修改商品" width="110">
      <template v-slot="scope">
        <div style="text-align: center">
          <el-icon>
            <Edit style="height: 25px; width: 25px;text-align: center;color: #409eff" @click="this.focus_commodity_item_id=scope.row.id; dialogVisible=true;">
            </Edit>
          </el-icon>
        </div>
        <el-dialog
          v-model="dialogVisible"
          title="修改商品"
          width="30%"
          append-to-body
        >
          <el-form :model="update_commodity_item" :rules="updateRules" ref="updateRules" label-width="120px">
            <el-form-item label="商品名" prop="item_name">
              <el-input v-model="update_commodity_item.item_name" placeholder="请输入修改后的商品名" />
            </el-form-item>
            <el-form-item label="商品价格" prop="item_price">
              <el-input-number
                v-model="update_commodity_item.item_price"
                :precision="2"
                :min="0.00"
                placeholder="请输入修改后的商品价格"
              />
            </el-form-item>
          </el-form>
          <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="dialogVisible = false;updateCommodityItem();">
          确定
        </el-button>
      </span>
          </template>
        </el-dialog>
      </template>
    </el-table-column>
    <el-table-column label="下架商品" width="110">
      <template v-slot="scope">
        <div style="text-align: center">
          <el-icon>
            <Delete
              style="height: 25px; width: 25px; text-align: center; color: #409eff"
              @click="showConfirmationDialog(scope.row.id)"
            >
            </Delete>
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
    :total="seller_commodity_item.length"
    position="bottom"
    background
  />
</template>

<script>

import {Edit, Plus,Delete} from "@element-plus/icons-vue";

export default {
  name: "seller_commodity_item",
  components: {Edit, Plus,Delete},
  data() {
    return {
      seller_commodity_item: localStorage.getItem('seller_commodity_item')
        ? JSON.parse(localStorage.getItem('seller_commodity_item'))
        : [],
      seller_data: localStorage.getItem('seller_data')
        ? JSON.parse(localStorage.getItem('seller_data'))
        : {},
      commodity_price_history: localStorage.getItem('commodity_price_history')
        ? JSON.parse(localStorage.getItem('commodity_price_history'))
        : [],
      focus_commodity_item_id:0,
      showPrice:false,
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
      update_commodity_item:{
        commodity_item_id:-1,
        item_name:null,
        price:0,
      },
      updateRules: {
        item_name: [
          { required: true, message: '请输入商品名', trigger: 'blur' },
          { min: 2, max: 20, message: "长度在 2 到 20 个字符", trigger: "blur" }
        ],
        item_price: [
          { required: true, message: '请输入商品价格', trigger: 'blur' },
          { type: 'number', message: '请输入有效的数字', trigger: 'blur' },
          { validator: this.validatePrice, trigger: 'blur' },
        ],
      },
      //page
      currentPage: 1,
      pageSize: 10,
      //drawer
      drawer: false,
      innerDrawer: false,
      //dialog
      dialogVisible:false,
      //key
      table_key1:'',
      table_key2:'',
    }
  },
  computed: {
    currentTableData() {
      const start = (this.currentPage - 1) * this.pageSize;
      const end = start + this.pageSize;
      return this.seller_commodity_item.slice(start, end);
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
    if(localStorage.getItem('seller_commodity_item') === null){
      this.findAll()
      this.commodity_price_history=[]
    }
  },
  methods: {
    handleCurrentChange(val) {
      this.currentPage = val;
    },
    findAll() {
      this.request.get("/commodity/all").then((res) => {
        if (res.status === 200) {
          localStorage.setItem("seller_commodity_item", JSON.stringify(res.data));
          this.table_key1 = Math.random()
        } else {
          this.$message.error(res.message);
        }
      }).catch(error => {
        this.$message.error(error.response.data.message);
      });
    },
    updateCommodityItem(){
      this.$refs["updateRules"].validate((valid) => {
        if (valid) {
          this.update_commodity_item.commodity_item_id=this.focus_commodity_item_id;
          this.request.put("/commodity/item",this.update_commodity_item).then((res) => {
            if (res.status === 200) {
              this.$message.success("设置成功")
              this.table_key1 = Math.random()
            } else {
              this.$message.error(res.message);
            }
          }).catch(error => {
            this.$message.error(error.response.data.message);
          });
        }
      });
    },
    validatePrice(rule, value, callback) {
      // 自定义价格验证规则示例
      if (value <= 0) {
        callback(new Error('商品价格必须大于零'));
      } else {
        callback();
      }
    },
    deleteCommodityItem(){
      this.request.delete("/commodity/item/"+this.focus_commodity_item_id).then((res) => {
        if (res.status === 200) {
          this.$message.success("删除成功")
          this.table_key1 = Math.random()
          location.reload();
        } else {
          this.$message.error(res.message);
        }
      }).catch(error => {
        this.$message.error(error.response.data.message);
      });
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
    showConfirmationDialog(itemId) {
      this.$confirm('确定要删除这个商品吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(() => {
          // 用户点击了确定按钮，可以在这里执行删除商品的操作
          this.focus_commodity_item_id = itemId;
          this.deleteCommodityItem();
          location.reload();
        })
        .catch(() => {
          // 用户点击了取消按钮，不执行任何操作
        });
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
  },
}
</script>

<style>
.el-table .lowest-price-row {
  background-color: #fa0404; /* 设置最低价格行的背景色 */
  /* 可以根据需要添加其他样式 */
}
</style>