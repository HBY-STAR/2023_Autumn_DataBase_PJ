<template>
  <div style="display: flex">
    <div style="width: 950px;">
      <el-table :data="currentTableData1" border height="550px" show-empty style="width: 950px">
        <el-table-column label="商品ID" prop="id" width="100"/>
        <el-table-column label="默认名称" prop="default_name" width="150"/>
        <el-table-column label="种类" prop="type" width="150"/>
        <el-table-column label="生产日期" prop="produce_at" width="150"/>
        <el-table-column label="生产地址" prop="produce_address" width="220"/>
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
              title="修改商品信息"
              width="30%"
            >
              <el-form :model="update_commodity" :rules="update_rules" ref="update_rules" label-width="120px">
                <el-form-item  prop="default_name">
                  <el-input v-model="update_commodity.default_name" placeholder="请输入商品默认名称" />
                </el-form-item>
                <el-form-item  prop="type">
                  <el-input v-model="update_commodity.type" placeholder="请输入商品种类" />
                </el-form-item>
                <el-form-item   prop="produce_at">
                  <el-date-picker
                    v-model="update_commodity.produce_at"
                    type="datetime"
                    placeholder="请选择生产时间"
                  />
                </el-form-item>
                <el-form-item   prop="produce_address">
                  <el-input v-model="update_commodity.produce_address" placeholder="请输入生产地址" />
                </el-form-item>
              </el-form>
              <template #footer>
                <span class="dialog-footer">
                  <el-button @click="dialogVisible = false">取消</el-button>
                  <el-button type="primary" @click="dialogVisible = false;updateCommodity();">
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
        :total="all_commodity_data.length"
        position="bottom"
        background
      />
    </div>
    <div>
      <el-card style="margin-left: 20px; width: 300px">
        <div style="margin: 10px 0; text-align: center; font-size: 24px; color: cornflowerblue"><b>新建商品</b></div>
        <div style="height: 20px"></div>
        <el-form :model="add_commodity" :rules="update_rules" ref="update_rules">
          <el-form-item  prop="default_name">
            <el-input v-model="add_commodity.default_name" placeholder="请输入商品默认名称" />
          </el-form-item>
          <el-form-item  prop="type">
            <el-input v-model="add_commodity.type" placeholder="请输入商品种类" />
          </el-form-item>
          <el-form-item   prop="produce_at">
            <el-date-picker
              v-model="add_commodity.produce_at"
              type="datetime"
              placeholder="请选择生产时间"
            />
          </el-form-item>
          <el-form-item   prop="produce_address">
            <el-input v-model="add_commodity.produce_address" placeholder="请输入生产地址" />
          </el-form-item>
          <el-form-item>
            <el-button
              style="margin-top: 20px"
              type="primary"
              size="default"
              autocomplete="off"
              @click="addCommodity()">确定
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
  name: "admin_commodity_data",
  components:{Edit,Delete},
  data() {
    return {
      all_commodity_data: localStorage.getItem('all_commodity_data')
        ? JSON.parse(localStorage.getItem('all_commodity_data'))
        : [],
      focus_id:0,
      //page1
      currentPage1: 1,
      pageSize1:20,
      //dialog
      dialogVisible:false,
      //add
      add_commodity:{
        default_name:null,
        type:null,
        produce_at:null,
        produce_address:null,
      },
      //update
      update_commodity:{
        id:null,
        default_name:null,
        type:null,
        produce_at:null,
        produce_address:null,
      },
      //rules
      update_rules:{
        default_name:[
          { required: true, message: '不能为空', trigger: 'blur' },
        ],
        type:[
          { required: true, message: '不能为空', trigger: 'blur' },
        ],
        produce_at:[
          { required: true, message: '不能为空', trigger: 'blur' },
        ],
        produce_address:[
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
      return this.all_commodity_data.slice(start, end);
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
      this.request.get("/commodities/all").then((res) => {
        if (res.code === "200") {
          localStorage.setItem("all_commodity_data", JSON.stringify(res.data));
        } else {
          this.$message.error(res.message);
        }
      });
    },
    //add
    addCommodity(){
      this.$refs["update_rules"].validate((valid) => {
        if (valid) {
          this.request.post("/commodities",this.add_commodity).then((res) => {
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
    updateCommodity(){
      this.$refs["update_rules"].validate((valid) => {
        if (valid) {
          this.request.put("/commodities",this.update_commodity).then((res) => {
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
    deleteCommodity(){
      this.request.delete("/commodities/"+this.focus_id).then((res) => {
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
        this.deleteCommodity();
      }).catch(() => {
      });
    },
  }
}
</script>