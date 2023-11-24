<template>
  <div style="margin-top: 10px">
    <el-form :model="shop" :rules="rules" ref="shopForm">
      <el-form-item prop="number">
        <el-input
            v-model="shop.number"
            placeholder="请输入要充值的金额"
        ></el-input>
      </el-form-item>
    </el-form>
    <div style="margin-top: 15px">
      <el-button autocomplete="off" type="primary" @click="rechargeShopAccount"
      >确定
      </el-button
      >
      <el-button
          autocomplete="off"
          style="position: absolute;right: 0;"
          type="warning"
          @click="$router.push('/seller_shop_information')"
      >取消
      </el-button
      >
    </div>
  </div>
</template>

<script>

export default {
  name: "recharge_account",
  data() {
    return {
      shop:{},
      rules:{
        number: [
          { required: true, message: '充值金额为空', trigger: 'blur' },
          { pattern: /^([1-9][0-9]{2,}.\d*)|([1-9][0-9]{2,})$/,message: '单次充值不小于100元',trigger: 'blur'},
        ],
      },
      user_data: localStorage.getItem('user_data')
          ? JSON.parse(localStorage.getItem('user_data'))
          : {},
      shop_data: localStorage.getItem('shop_data')
          ? JSON.parse(localStorage.getItem('shop_data'))
          :{},
    }
  },
  methods: {
    rechargeShopAccount() {
      this.$refs['shopForm'].validate((valid) => {
        if (valid) {
          this.request.post('/shop-account/recharge_shop_account_'+this.shop_data.shopId, this.shop.number).then((res) => {
            if (res.code === '200') {
              localStorage.setItem('shop_account', JSON.stringify(res.data)) // 存储用户信息到浏览器
              this.$message.success("充值成功")
            } else {
              this.$message.error(res.msg)
            }
          })
        }
      })
    },
  }
}
</script>