<template>
  <div style="display: flex">
    <div style="width: 950px;">
      <el-table :data="currentTableData1" border height="550px" show-empty style="width: 950px">
        <el-table-column label="用户ID" prop="id" width="100"/>
        <el-table-column label="用户名" prop="username" width="150"/>
        <el-table-column label="密码" prop="password" width="150"/>
        <el-table-column label="年龄" prop="age" width="60"/>
        <el-table-column label="性别" prop="gender" width="60">
          <template v-slot="scope">
            {{ scope.row.gender ? '男' : '女' }}
          </template>
        </el-table-column>

        <el-table-column label="邮箱" prop="email" width="150"/>
        <el-table-column label="电话" prop="phone" width="100"/>
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
              <div style="margin-bottom: 30px; text-align: center; font-size: 24px; color: cornflowerblue"><b>修改用户信息</b></div>
              <el-form :model="update_user" :rules="update_rules" ref="update_rules">
                <el-form-item  prop="username">
                  <el-input v-model="update_user.username" placeholder="请输入用户名" />
                </el-form-item>
                <el-form-item  prop="password">
                  <el-input v-model="update_user.password" placeholder="请输入用户密码" />
                </el-form-item>
                <el-form-item  prop="age">
                  <el-input v-model="update_user.age" placeholder="请输入用户年龄" />
                </el-form-item>
                <el-form-item prop="email">
                  <el-input v-model="update_user.email" placeholder="请输入用户邮箱" />
                </el-form-item>
                <el-form-item prop="phone">
                  <el-input v-model="update_user.phone" placeholder="请输入用户电话" />
                </el-form-item>
                <el-form-item prop="gender">
                  <el-radio-group v-model="update_user.gender">
                    <el-radio label=true border>男</el-radio>
                    <el-radio label=false border>女</el-radio>
                  </el-radio-group>
                </el-form-item>
              </el-form>
              <template #footer>
                <span class="dialog-footer">
                  <el-button @click="dialogVisible = false">取消</el-button>
                  <el-button type="primary" @click="dialogVisible = false;updateUser();">
                    确定
                  </el-button>
                </span>
              </template>
            </el-dialog>
          </template>
        </el-table-column>
        <el-table-column label="删除用户" width="90">
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
        :total="all_user_data.length"
        position="bottom"
        background
      />
    </div>
    <div>
      <el-card style="margin-left: 20px; width: 300px">
        <div style="margin: 10px 0; text-align: center; font-size: 24px; color: cornflowerblue"><b>新建用户</b></div>
        <div style="height: 20px"></div>
        <el-form :model="add_user" :rules="update_rules" ref="update_rules">
          <el-form-item  prop="username">
            <el-input v-model="add_user.username" placeholder="请输入用户名" />
          </el-form-item>
          <el-form-item  prop="password">
            <el-input v-model="add_user.password" placeholder="请输入用户密码" />
          </el-form-item>
          <el-form-item  prop="age">
            <el-input v-model="add_user.age" placeholder="请输入用户年龄"/>
          </el-form-item>
          <el-form-item prop="email">
            <el-input v-model="add_user.email" placeholder="请输入用户邮箱" />
          </el-form-item>
          <el-form-item prop="phone">
            <el-input v-model="add_user.phone" placeholder="请输入用户电话" />
          </el-form-item>
          <el-form-item prop="gender">
            <el-radio-group v-model="add_user.gender">
              <el-radio label=true border>男</el-radio>
              <el-radio label=false border>女</el-radio>
            </el-radio-group>
          </el-form-item>
          <el-form-item>
            <el-button
              style="margin-top: 20px"
              type="primary"
              size="default"
              autocomplete="off"
              @click="addUser()">确定
            </el-button>
          </el-form-item>
        </el-form>
      </el-card>
    </div>
  </div>
</template>

<script>

import { Edit, Delete,} from "@element-plus/icons-vue";

export default {
  name: "admin_user_data",
  components:{Edit,Delete},
  data() {
    return {
      all_user_data: localStorage.getItem('all_user_data')
        ? JSON.parse(localStorage.getItem('all_user_data'))
        : [],
      focus_id:0,
      //page1
      currentPage1: 1,
      pageSize1:20,
      //dialog
      dialogVisible:false,
      //add
      add_user:{
        username:null,
        password:null,
        age:null,
        gender:null,
        email:null,
        phone:null,
      },
      //update
      update_user:{
        id:null,
        username:null,
        password:null,
        age:null,
        gender:null,
        email:null,
        phone:null,
      },
      //rules
      update_rules:{
        username:[
          { required: true, message: '不能为空', trigger: 'blur' },
        ],
        password:[
          { required: true, message: '不能为空', trigger: 'blur' },
        ],
        age:[
          { required: true, message: '不能为空', trigger: 'blur' },
        ],
        gender:[
          { required: true, message: '不能为空', trigger: 'blur' },
        ],
        email:[
          { required: true, message: '不能为空', trigger: 'blur' },
        ],
        phone:[
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
      return this.all_user_data.slice(start, end);
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
      this.request.get("/users/all").then((res) => {
        if (res.status === 200) {
          localStorage.setItem("all_user_data", JSON.stringify(res.data));
        } else {
          this.$message.error(res.message);
        }
      });
    },
    //add
    addUser(){
      this.$refs["update_rules"].validate((valid) => {
        if (valid) {
          this.add_user.age = parseInt(this.add_user.age)
          this.add_user.gender = Boolean(this.add_user.gender)
          this.request.post("/users",this.add_user).then((res) => {
            if (res.status === 200) {
              this.$message.success("新建成功")
            } else {
              this.$message.error(res.message);
            }
          });
        }
      });
    },
    //update
    updateUser(){
      this.$refs["update_rules"].validate((valid) => {
        if (valid) {
          this.update_user.id = this.focus_id
          this.update_user.age = parseInt(this.update_user.age)
          this.update_user.gender = Boolean(this.update_user.gender)
          this.request.put("/users",this.update_user).then((res) => {
            if (res.status === 200) {
              this.$message.success("修改成功")
            } else {
              this.$message.error(res.message);
            }
          });
        }
      });
    },
    //delete
    deleteUser(){
      this.request.delete("/users/"+this.focus_id).then((res) => {
        if (res.status === 200) {
          this.$message.success("删除成功")
        } else {
          this.$message.error(res.message);
        }
      });
    },
    showDeleteConfirm(row) {
      this.$confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      }).then(() => {
        this.focus_id = row.id;
        this.deleteUser();
      }).catch(() => {
      });
    },
    handleDialogClose(){
      this.dialogVisible=false
    }
  }
}
</script>