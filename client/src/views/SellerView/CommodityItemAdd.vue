<template>
  <div style="display: flex">
    <div style="width: 900px;">
      <el-table :data="currentTableData1" border height="300px" show-empty style="width: 900px">
        <el-table-column label="可发布商品ID" prop="id" width="120"/>
        <el-table-column label="默认名称" prop="default_name" width="150"/>
        <el-table-column label="种类" prop="category" width="150"/>
        <el-table-column label="生产日期" prop="produce_at" width="200"/>
        <el-table-column label="生产地址" prop="produce_address" width="280"/>
      </el-table>
      <el-pagination
        style="margin-top: 10px"
        small
        @current-change="handleCurrentChange1"
        :current-page="currentPage1"
        :page-size="pageSize1"
        layout="prev, pager, next"
        :total="commodities_all.length"
        position="bottom"
      />
      <el-table :data="currentTableData2" border height="180px" show-empty style="width: 900px; margin-top: 30px">
        <el-table-column label="可发布平台ID" prop="id" width="150"/>
        <el-table-column label="平台名" prop="name" width="150"/>
        <el-table-column label="所在国家" prop="country" width="150"/>
        <el-table-column label="url" prop="url" width="450"/>
      </el-table>
      <el-pagination
        style="margin-top: 10px"
        small
        @current-change="handleCurrentChange2"
        :current-page="currentPage2"
        :page-size="pageSize2"
        layout="prev, pager, next"
        :total="platform_all.length"
        position="bottom"
      />
    </div>
    <div>
      <el-card style="margin-left: 40px; width: 350px">
        <div style="margin: 10px 0; text-align: center; font-size: 24px; color: cornflowerblue"><b>发布商品</b></div>
        <div style="height: 20px"></div>
        <el-form :model="addCommodity" :rules="addRules" ref="addRules">
          <el-form-item prop="commodity_id">
            <el-input
              style="margin-top: 20px"
              size="large"
              prefix-icon="Goods"
              v-model="addCommodity.commodity_id"
              placeholder="请输入商品ID"
            ></el-input>
          </el-form-item>
          <el-form-item prop="platform_id">
            <el-input
              style="margin-top: 20px"
              size="large"
              prefix-icon="Place"
              v-model="addCommodity.platform_id"
              placeholder="请输入平台ID"
            ></el-input>
          </el-form-item>
          <el-form-item prop="item_name">
            <el-input
              style="margin-top: 20px"
              size="large"
              prefix-icon="Checked"
              v-model="addCommodity.item_name"
              placeholder="请输入发布商品名"
            ></el-input>
          </el-form-item>
          <el-form-item prop="price">
            <el-input
              style="margin-top: 20px"
              size="large"
              prefix-icon="Money"
              v-model="addCommodity.price"
              placeholder="请输入发布商品价格"
            />
          </el-form-item>
          <el-form-item>
            <el-button
              style="margin-top: 20px"
              type="primary"
              size="default"
              autocomplete="off"
              @click="addCommodityItem()">发布
            </el-button>
          </el-form-item>
        </el-form>
      </el-card>
    </div>
  </div>
</template>

<script>

export default {
  name: "seller_add_item",
  data() {
    return {
      commodities_all: localStorage.getItem('commodities_all')
        ? JSON.parse(localStorage.getItem('commodities_all'))
        : [],
      platform_all:localStorage.getItem('platform_all')
        ? JSON.parse(localStorage.getItem('platform_all'))
        : [],
      focus_commodity_id:0,
      //page1
      currentPage1: 1,
      pageSize1:10,
      //page2
      currentPage2: 1,
      pageSize2:5,
      //add
      addCommodity:{
        commodity_id: null,
        platform_id: null,
        item_name: null,
        price: null,
      },
      addRules:{
        commodity_id: [
          { required: true, message: "商品ID不能为空", trigger: "blur" },
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
    currentTableData1() {
      const start = (this.currentPage1 - 1) * this.pageSize1;
      const end = start + this.pageSize1;
      return this.commodities_all.slice(start, end);
    },
    currentTableData2() {
      const start = (this.currentPage2 - 1) * this.pageSize2;
      const end = start + this.pageSize2;
      return this.platform_all.slice(start, end);
    },
  },
  created() {
    this.findAll()
  },
  methods: {
    handleCurrentChange1(val) {
      this.currentPage1 = val;
    },
    handleCurrentChange2(val) {
      this.currentPage2 = val;
    },
    validatePrice(rule, value, callback) {
      // 自定义价格验证规则示例
      if (value <= 0) {
        callback(new Error('商品价格必须大于零'));
      } else {
        callback();
      }
    },
    findAll() {
      this.request.get("/commodities/all").then((res) => {
        if (res.status === 200) {
          localStorage.setItem("commodities_all", JSON.stringify(res.data));
        } else {
          this.$message.error(res.message);
        }
      });
      this.request.get("/platform/all").then((res) => {
        if (res.status === 200) {
          localStorage.setItem("platform_all", JSON.stringify(res.data));
        } else {
          this.$message.error(res.message);
        }
      });
    },
    addCommodityItem(){
      this.$refs["addRules"].validate((valid) => {
        if (valid) {
          this.addCommodity.price=parseFloat(this.addCommodity.price)
          this.addCommodity.commodity_id=parseInt(this.addCommodity.commodity_id)
          this.addCommodity.platform_id=parseInt(this.addCommodity.platform_id)
          this.request.post("/commodity/item",this.addCommodity).then((res) => {
            if (res.status === 200) {
              this.$message.success("发布成功")
              location.reload();
            } else {
              this.$message.error(res.message);
            }
          });
        }
      });
    },
  }
}
</script>