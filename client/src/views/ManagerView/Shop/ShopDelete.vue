<template>
  <el-container style="height: 500px">
    <el-scrollbar height="500px">
      <el-table :data="shop_delete_data" border="true" height="500px">
        <el-table-column label="商店序列号" prop="shopId" width="100"/>
        <el-table-column label="商店名" prop="shopName" width="90"/>
        <el-table-column label="商品种类" prop="productType" width="120"/>
        <el-table-column label="简历" prop="shopResume" width="180"/>
        <el-table-column label="备案地址" prop="shopAddress" width="120"/>
        <el-table-column label="注册资金" prop="shopRegisterCost" width="90"/>
        <el-table-column label="注册时间" prop="shopRegisterTime" width="160"/>
        <el-table-column label="商店分数" prop="shopScore" width="90"/>
        <el-table-column label="商店状态" prop="shopState" width="90"/>
        <el-table-column label="店主UID" prop="sellerUid" width="90"/>
        <el-table-column label="店主身份证号" prop="sellerIdNumber" width="180"/>
      </el-table>
    </el-scrollbar>
  </el-container>
  <el-card style="margin-top: 30px">
    <el-form ref="checkForm" :model="check_shop_request" :rules="check_rules" >
      <el-space :size="100" alignment="center" direction="horizontal">
        <el-form-item>
        <span style="font-size: medium">
          处理商店注销请求
        </span>
        </el-form-item>
        <el-form-item prop="shop_id">
          <el-input
              v-model="check_shop_request.shop_id"
              placeholder="请输入商店ID"
              prefix-icon="Edit"
              size="large"
          ></el-input>
        </el-form-item>
        <el-form-item prop="is_pass">
          <el-radio-group v-model="check_shop_request.is_pass">
            <el-radio border label="通过">通过</el-radio>
            <el-radio border label="不通过">不通过</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item style="text-align: right">
          <el-button autocomplete="off" size="medium" type="primary" @click="check_request"
          >确定
          </el-button
          >
        </el-form-item>
      </el-space>
    </el-form>
  </el-card>
</template>

<script>

export default {
  name: "shop_delete_data",
  data() {
    return {
      check_shop_request: {},
      check_rules: {
        shop_id: [
          {required: true, message: '商店ID不能为空', trigger: 'blur'},
          {pattern: /^\d{5,}$/, message: '商店ID应为至少5位数字', trigger: 'blur'},
        ],
        is_pass: [
          {required: true, message: '请选择是否通过该请求', trigger: 'change'}
        ],
      },
      shop_delete_data: localStorage.getItem('shop_delete_data')
          ? JSON.parse(localStorage.getItem('shop_delete_data'))
          : {},
    }
  },
  created() {
    this.findAll()
  },
  methods: {
    findAll() {
      this.request.get('/shop-data/find_delete_shop').then((res) => {
        if (res.code === '200') {
          localStorage.setItem('shop_delete_data', JSON.stringify(res.data))
        } else {
          this.$message.error(res.msg)
        }
      })
    },
    check_request() {
      this.$refs['checkForm'].validate((valid) => {
        if (valid) {
          this.request.post('/shop-data/check_shop_delete_'+this.check_shop_request.shop_id, this.check_shop_request.is_pass).then((res) => {
            if (res.code === '200') {
              this.$message.success("该请求已处理")
            } else {
              this.$message.error(res.msg)
            }
          })
        }
      })
    }
  }
}
</script>