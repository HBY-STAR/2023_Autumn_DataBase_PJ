<template>
  <el-container style="height: 500px">
    <el-scrollbar height="500px">
      <el-table :data="order_request_data" border="true" height="500px">
        <el-table-column label="订单序列号" prop="recordId" width="100"/>
        <el-table-column label="用户ID" prop="uid" width="200"/>
        <el-table-column label="商品ID" prop="productId" width="200"/>
        <el-table-column label="商品价格" prop="price" width="180"/>
        <el-table-column label="商品状态" prop="state" width="200"/>
        <el-table-column label="创建时间" prop="time" width="200"/>
      </el-table>
    </el-scrollbar>
  </el-container>
  <el-card style="margin-top: 30px;width: 1100px">
    <el-form ref="checkForm" :model="check_order_request" :rules="check_rules">
      <el-space :size="100" alignment="center" direction="horizontal">
        <el-form-item>
        <span style="font-size: medium">
          处理订单
        </span>
        </el-form-item>
        <el-form-item prop="productId">
          <el-input
              v-model="check_order_request.orderId"
              placeholder="请输入订单ID"
              prefix-icon="Edit"
              size="large"
          ></el-input>
        </el-form-item>
        <el-form-item prop="isPass">
          <el-radio-group v-model="check_order_request.isPass">
            <el-radio border label="通过">通过</el-radio>
            <el-radio border label="不通过">不通过</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item style="text-align: right">
          <el-button autocomplete="off" type="primary" @click="check_request"
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
  name: "order_request_data",
  data() {
    return {
      check_order_request: {},
      check_rules: {
        orderId: [
          {required: true, message: '订单号不能为空', trigger: 'blur'},
          {pattern: /^\d{5,}$/, message: '订单号应为至少5位数字', trigger: 'blur'},
        ],
        isPass: [
          {required: true, message: '请选择是否通过该请求', trigger: 'change'}
        ],
      },
      order_request_data: localStorage.getItem('order_request_data')
          ? JSON.parse(localStorage.getItem('order_request_data'))
          : {},
      shop_data: localStorage.getItem('shop_data')
          ? JSON.parse(localStorage.getItem('shop_data'))
          :{},
      user_data: localStorage.getItem('user_data')
          ? JSON.parse(localStorage.getItem('user_data'))
          : {},
    }
  },
  created() {
    this.findShopId()
    this.findAll()
  },
  methods: {
    findAll() {
      this.request.get('/all-trade-record/find_order_by_'+this.shop_data.shopId).then((res) => {
        if (res.code === '200') {
          localStorage.setItem('order_request_data', JSON.stringify(res.data))
        } else {
          this.$message.error(res.msg)
        }
      })
    },
    findShopId(){
      this.request.get('/shop-data/find_by_uid_' + this.user_data.uid).then((res) => {
        if (res.code === '200') {
          localStorage.setItem('shop_data', JSON.stringify(res.data)) // 存储用户信息到浏览器
        } else {
          this.$message.error(res.msg)
        }
      })
    },
    check_request() {
      this.$refs['checkForm'].validate((valid) => {
        if (valid) {
          this.request.post('/all-trade-record/check_order_request', this.check_order_request).then((res) => {
            if (res.code === '200') {
              this.$message.success("该订单已处理")
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