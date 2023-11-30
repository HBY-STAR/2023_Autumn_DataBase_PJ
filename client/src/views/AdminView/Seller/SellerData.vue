<template>
  <div style="display: flex">
    <div style="width: 950px;">
      <el-table :data="currentTableData1" border height="550px" show-empty style="width: 950px">
        <el-table-column label="商家ID" prop="id" width="100"/>
        <el-table-column label="商家名" prop="username" width="150"/>
        <el-table-column label="密码" prop="password" width="150"/>
        <el-table-column label="邮箱" prop="email" width="150"/>
        <el-table-column label="地址" prop="address" width="220"/>
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
              title="修改商家信息"
              width="30%"
            >
              <el-form :model="update_seller" :rules="update_rules" ref="update_rules" label-width="120px">
                <el-form-item  prop="username">
                  <el-input v-model="update_seller.username" placeholder="请输入商家名" />
                </el-form-item>
                <el-form-item  prop="password">
                  <el-input v-model="update_seller.password" placeholder="请输入商家密码" />
                </el-form-item>
                <el-form-item prop="email">
                  <el-input v-model="update_seller.email" placeholder="请输入商家邮箱" />
                </el-form-item>
                <el-form-item prop="address">
                  <el-input v-model="update_seller.address" placeholder="请输入商家地址" />
                </el-form-item>
              </el-form>
              <template #footer>
                <span class="dialog-footer">
                  <el-button @click="dialogVisible = false">取消</el-button>
                  <el-button type="primary" @click="dialogVisible = false;updateSeller();">
                    确定
                  </el-button>
                </span>
              </template>
            </el-dialog>
          </template>
        </el-table-column>
        <el-table-column label="删除商家" width="90">
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
        :total="all_seller_data.length"
        position="bottom"
        background
      />
    </div>
    <div>
      <el-card style="margin-left: 20px; width: 300px">
        <div style="margin: 10px 0; text-align: center; font-size: 24px; color: cornflowerblue"><b>新建商家</b></div>
        <div style="height: 20px"></div>
        <el-form :model="add_seller" :rules="update_rules" ref="update_rules">
          <el-form-item  prop="username">
            <el-input v-model="add_seller.username" placeholder="请输入商家名" />
          </el-form-item>
          <el-form-item  prop="password">
            <el-input v-model="add_seller.password" placeholder="请输入商家密码" />
          </el-form-item>
          <el-form-item prop="email">
            <el-input v-model="add_seller.email" placeholder="请输入商家邮箱" />
          </el-form-item>
          <el-form-item prop="address">
            <el-input v-model="add_seller.address" placeholder="请输入商家地址" />
          </el-form-item>
          <el-form-item>
            <el-button
              style="margin-top: 20px"
              type="primary"
              size="default"
              autocomplete="off"
              @click="addSeller()">确定
            </el-button>
          </el-form-item>
        </el-form>
      </el-card>
    </div>
  </div>
</template>

<script>

import { Edit, Delete, MessageBox } from "@element-plus/icons-vue";

export default {
  name: "admin_seller_data",
  components:{Edit,Delete},
  data() {
    return {
      all_seller_data: localStorage.getItem('all_seller_data')
        ? JSON.parse(localStorage.getItem('all_seller_data'))
        : [],
      focus_id:0,
      //page1
      currentPage1: 1,
      pageSize1:20,
      //dialog
      dialogVisible:false,
      //add
      add_seller:{
        username:null,
        password:null,
        email:null,
        address:null,
      },
      //update
      update_seller:{
        id:null,
        username:null,
        password:null,
        email:null,
        address:null,
      },
      //rules
      update_rules:{
        username:[
          { required: true, message: '不能为空', trigger: 'blur' },
        ],
        password:[
          { required: true, message: '不能为空', trigger: 'blur' },
        ],
        email:[
          { required: true, message: '不能为空', trigger: 'blur' },
        ],
        address:[
          { required: true, message: '不能为空', trigger: 'blur' },
        ],
      }
    }
  },
  computed: {
    //page
    currentTableData1() {
      const start = (this.currentPage1 - 1) * this.pageSize1;
      const end = start + this.pageSize1;
      return this.all_seller_data.slice(start, end);
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
      this.request.get("/sellers/data").then((res) => {
        if (res.code === "200") {
          localStorage.setItem("all_seller_data", JSON.stringify(res.data));
        } else {
          this.$message.error(res.message);
        }
      });
    },
    //add
    addSeller(){
      this.$refs["update_rules"].validate((valid) => {
        if (valid) {
          this.request.post("/sellers",this.add_seller).then((res) => {
            if (res.code === "200") {
              this.$message.success("新建成功")
            } else {
              this.$message.error(res.message);
            }
          });
        }
      });
    },
    //update
    updateSeller(){
      this.$refs["update_rules"].validate((valid) => {
        if (valid) {
          this.request.put("/sellers",this.update_seller).then((res) => {
            if (res.code === "200") {
              this.$message.success("修改成功")
            } else {
              this.$message.error(res.message);
            }
          });
        }
      });
    },
    //delete
    deleteSeller(){
      this.request.delete("/sellers/"+this.focus_id).then((res) => {
        if (res.code === "200") {
          this.$message.success("删除成功")
        } else {
          this.$message.error(res.message);
        }
      });
    },
    showDeleteConfirm(row) {
      MessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      }).then(() => {
        this.focus_id = row.id;
        this.deleteSeller();
      }).catch(() => {
      });
    },
  }
}
</script>