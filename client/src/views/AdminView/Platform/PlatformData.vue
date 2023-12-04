<template>
  <div style="display: flex">
    <div style="width: 950px;">
      <el-table :data="currentTableData1" border height="550px" show-empty style="width: 950px">
        <el-table-column label="平台ID" prop="id" width="150"/>
        <el-table-column label="平台名" prop="name" width="150"/>
        <el-table-column label="所在国家" prop="country" width="120"/>
        <el-table-column label="url" prop="url" width="350"/>
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
              title="修改平台信息"
              width="30%"
            >
              <el-form :model="update_platform" :rules="update_rules" ref="update_rules" label-width="120px">
                <el-form-item  prop="name">
                  <el-input v-model="update_platform.name" placeholder="请输入平台名" />
                </el-form-item>
                <el-form-item  prop="country">
                  <el-input v-model="update_platform.country" placeholder="请输入平台所属国家" />
                </el-form-item>
                <el-form-item   prop="url">
                  <el-input v-model="update_platform.url" placeholder="请输入平台URL" />
                </el-form-item>
              </el-form>
              <template #footer>
                <span class="dialog-footer">
                  <el-button @click="dialogVisible = false">取消</el-button>
                  <el-button type="primary" @click="dialogVisible = false;updatePlatform();">
                    确定
                  </el-button>
                </span>
              </template>
            </el-dialog>
          </template>
        </el-table-column>
        <el-table-column label="删除平台" width="90">
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
        :total="all_platform_data.length"
        position="bottom"
        background
      />
    </div>
    <div>
      <el-card style="margin-left: 20px; width: 300px">
        <div style="margin: 10px 0; text-align: center; font-size: 24px; color: cornflowerblue"><b>新建平台</b></div>
        <div style="height: 20px"></div>
        <el-form :model="add_platform" :rules="update_rules" ref="update_rules">
          <el-form-item  prop="name">
            <el-input v-model="add_platform.name" placeholder="请输入平台名" />
          </el-form-item>
          <el-form-item  prop="country">
            <el-input v-model="add_platform.country" placeholder="请输入平台所属国家" />
          </el-form-item>
          <el-form-item   prop="url">
            <el-input v-model="add_platform.url" placeholder="请输入平台URL" />
          </el-form-item>
          <el-form-item>
            <el-button
              style="margin-top: 20px"
              type="primary"
              size="default"
              autocomplete="off"
              @click="addPlatform()">确定
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
  name: "admin_platform_data",
  components:{Edit,Delete},
  data() {
    return {
      all_platform_data: localStorage.getItem('all_platform_data')
        ? JSON.parse(localStorage.getItem('all_platform_data'))
        : [],
      focus_id:0,
      //page1
      currentPage1: 1,
      pageSize1:20,
      //dialog
      dialogVisible:false,
      //add
      add_platform:{
        name:null,
        country:null,
        url:null,
      },
      //update
      update_platform:{
        id:null,
        name:null,
        country:null,
        url:null,
      },
      //rules
      update_rules:{
        name:[
          { required: true, message: '不能为空', trigger: 'blur' },
        ],
        country:[
          { required: true, message: '不能为空', trigger: 'blur' },
        ],
        url:[
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
      return this.all_platform_data.slice(start, end);
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
      this.request.get("/platform/all").then((res) => {
        if (res.status === 200) {
          localStorage.setItem("all_platform_data", JSON.stringify(res.data));
        } else {
          this.$message.error(res.message);
        }
      });
    },
    //add
    addPlatform(){
      this.$refs["update_rules"].validate((valid) => {
        if (valid) {
          this.request.post("/platform",this.add_platform).then((res) => {
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
    updatePlatform(){
      this.$refs["update_rules"].validate((valid) => {
        if (valid) {
          this.request.put("/platform",this.update_platform).then((res) => {
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
    deletePlatform(){
      this.request.delete("/platform/"+this.focus_id).then((res) => {
        if (res.status === 200) {
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
        this.deletePlatform();
      }).catch(() => {
      });
    },
  }
}
</script>