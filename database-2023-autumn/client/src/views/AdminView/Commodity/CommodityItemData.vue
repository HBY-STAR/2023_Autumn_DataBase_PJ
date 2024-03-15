<template>
  <div style="display: flex">
    <div style="width: 950px;">
      <el-table :data="currentTableData1" border height="550px" show-empty style="width: 950px">
        <el-table-column label="商品序号" prop="id" width="150"/>
        <el-table-column label="商品名" prop="item_name" width="200"/>
        <el-table-column label="当前价格" prop="price" width="120"/>
        <el-table-column label="商品ID" prop="commodity_id" width="100"/>
        <el-table-column label="商家ID" prop="seller_id" width="100"/>
        <el-table-column label="平台ID" prop="platform_id" width="100"/>
        <el-table-column label="修改信息" width="90">
          <template v-slot="scope">
            <div style="text-align: center">
              <el-icon>
                <Edit style="height: 25px; width: 25px;text-align: center;color: #409eff" @click="this.focus_id=scope.row.id;dialogVisible=true;">
                </Edit>
              </el-icon>
            </div>
            <el-dialog
              v-model="dialogVisible"
              width="25%"
              :append-to-body="true"
              :before-close="handleDialogClose"
            >
              <div style="margin-bottom: 30px; text-align: center; font-size: 24px; color: cornflowerblue"><b>修改商品信息</b></div>
              <el-form :model="update_commodity_item" :rules="update_rules" ref="update_rules">
                <el-form-item prop="commodity_id">
                  <el-input
                    v-model="update_commodity_item.commodity_id"
                    placeholder="请输入商品ID"
                  ></el-input>
                </el-form-item>
                <el-form-item prop="seller_id">
                  <el-input
                    v-model="update_commodity_item.seller_id"
                    placeholder="请输入商家ID"
                  ></el-input>
                </el-form-item>
                <el-form-item prop="platform_id">
                  <el-input
                    v-model="update_commodity_item.platform_id"
                    placeholder="请输入平台ID"
                  ></el-input>
                </el-form-item>
                <el-form-item prop="item_name">
                  <el-input
                    v-model="update_commodity_item.item_name"
                    placeholder="请输入发布商品名"
                  ></el-input>
                </el-form-item>
                <el-form-item prop="price">
                  <el-input
                    v-model="update_commodity_item.price"
                    placeholder="请输入发布商品价格"
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
        <el-table-column label="删除商品" width="90">
          <template v-slot="scope">
            <div style="text-align: center">
              <el-icon>
                <Delete style="height: 25px; width: 25px; text-align: center; color: #409eff" @click="showDeleteConfirm(scope.row)">
                </Delete>
              </el-icon>
            </div>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        style="margin-top: 10px"
        small
        @current-change="handleCurrentChange1"
        :current-page="currentPage1"
        :page-size="pageSize1"
        layout="total, prev, pager, next, jumper"
        :total="all_commodity_item_data.length"
        position="bottom"
        background
      />
    </div>
    <div>
      <el-card style="margin-left: 20px; width: 300px">
        <div style="margin: 10px 0; text-align: center; font-size: 24px; color: cornflowerblue"><b>新建商品</b></div>
        <div style="height: 20px"></div>
        <el-form :model="add_commodity_item" :rules="update_rules" ref="update_rules">
          <el-form-item prop="commodity_id">
            <el-input
              v-model="add_commodity_item.commodity_id"
              placeholder="请输入商品ID"
            ></el-input>
          </el-form-item>
          <el-form-item prop="seller_id">
            <el-input
              v-model="add_commodity_item.seller_id"
              placeholder="请输入商家ID"
            ></el-input>
          </el-form-item>
          <el-form-item prop="platform_id">
            <el-input
              v-model="add_commodity_item.platform_id"
              placeholder="请输入平台ID"
            ></el-input>
          </el-form-item>
          <el-form-item prop="item_name">
            <el-input
              v-model="add_commodity_item.item_name"
              placeholder="请输入发布商品名"
            ></el-input>
          </el-form-item>
          <el-form-item prop="price">
            <el-input
              v-model="add_commodity_item.price"
              placeholder="请输入发布商品价格"
            />
          </el-form-item>
          <el-form-item>
            <el-button
              style="margin-top: 20px"
              type="primary"
              size="default"
              autocomplete="off"
              @click="addCommodityItem()">确定
            </el-button>
          </el-form-item>
        </el-form>
      </el-card>
    </div>
  </div>
</template>

<script>

import { Edit, Delete} from "@element-plus/icons-vue";

export default {
  name: "admin_commodity_item_data",
  components:{Edit,Delete},
  data() {
    return {
      all_commodity_item_data: localStorage.getItem('all_commodity_item_data')
        ? JSON.parse(localStorage.getItem('all_commodity_item_data'))
        : [],
      focus_id:0,
      //page1
      currentPage1: 1,
      pageSize1:20,
      //dialog
      dialogVisible:false,
      //add
      add_commodity_item:{
        seller_id:null,
        commodity_id: null,
        platform_id: null,
        item_name: null,
        price: null,
      },
      //update
      update_commodity_item:{
        commodity_item_id:null,
        seller_id:null,
        commodity_id: null,
        platform_id: null,
        item_name: null,
        price: null,
      },
      //rules
      update_rules:{
        commodity_id: [
          { required: true, message: "商品ID不能为空", trigger: "blur" },
        ],
        seller_id:[
          { required: true, message: "商家ID不能为空", trigger: "blur" },
        ],
        platform_id : [
          { required: true, message: "平台ID不能为空", trigger: "blur" },
        ],
        item_name: [
          { required: true, message: "商品名不能为空", trigger: "blur" },
          { min: 2, max: 20, message: "长度在 2 到 20 个字符", trigger: "blur" }
        ],
        price: [
          { required: true, message: '商品价格不能为空', trigger: 'blur' },
          { validator: this.validatePrice, trigger: 'blur' },
        ],
      }
    }
  },
  computed: {
    //page
    currentTableData1() {
      const start = (this.currentPage1 - 1) * this.pageSize1;
      const end = start + this.pageSize1;
      return this.all_commodity_item_data.slice(start, end);
    },
  },
  created() {
    this.findAll()
  },
  methods: {
    //page
    handleCurrentChange1(val) {
      this.currentPage1 = val;
    },
    //data
    findAll() {
      this.request.get("/commodity/all").then((res) => {
        if (res.status === 200) {
          localStorage.setItem("all_commodity_item_data", JSON.stringify(res.data));
        } else {
          this.$message.error(res.message);
        }
      }).catch(error => {
        // 这里处理通过拦截器输出的错误信息
        this.$message.error(error.response.data.message);
      });
    },
    //add
    addCommodityItem(){
      this.$refs["update_rules"].validate((valid) => {
        if (valid) {
          this.add_commodity_item.commodity_id=parseInt(this.add_commodity_item.commodity_id)
          this.add_commodity_item.seller_id=parseInt(this.add_commodity_item.seller_id)
          this.add_commodity_item.platform_id=parseInt(this.add_commodity_item.platform_id)
          this.add_commodity_item.price=parseFloat(this.add_commodity_item.price)
          this.request.post("/commodity/item",this.add_commodity_item).then((res) => {
            if (res.status === 200) {
              this.$message.success("新建成功")
            } else {
              this.$message.error(res.message);
            }
          }).catch(error => {
            // 这里处理通过拦截器输出的错误信息
            this.$message.error(error.response.data.message);
          });
        }
      });
    },
    //update
    updateCommodityItem(){
      this.$refs["update_rules"].validate((valid) => {
        if (valid) {
          this.update_commodity_item.commodity_item_id=this.focus_id
          this.update_commodity_item.commodity_id = parseInt(this.update_commodity_item.commodity_id)
          this.update_commodity_item.platform_id=parseInt(this.update_commodity_item.platform_id)
          this.update_commodity_item.price=parseFloat(this.update_commodity_item.price)
          this.update_commodity_item.seller_id=0
          this.request.put("/commodity/item",this.update_commodity_item).then((res) => {
            if (res.status === 200) {
              this.$message.success("修改成功")
            } else {
              this.$message.error(res.message);
            }
          }).catch(error => {
            // 这里处理通过拦截器输出的错误信息
            this.$message.error(error.response.data.message);
          });
        }
      });
    },
    //delete
    deleteCommodityItem(){
      this.request.delete("/commodity/item/"+this.focus_id).then((res) => {
        if (res.status === 200) {
          this.$message.success("删除成功")
        } else {
          this.$message.error(res.message);
        }
      }).catch(error => {
        // 这里处理通过拦截器输出的错误信息
        this.$message.error(error.response.data.message);
      });
    },
    showDeleteConfirm(row) {
      this.$confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      }).then(() => {
        this.focus_id = row.id;
        this.deleteCommodityItem();
      }).catch(() => {
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
    handleDialogClose(){
      this.dialogVisible=false
    },
  },


}
</script>