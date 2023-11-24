<template>
  <div>
    <el-scrollbar height="500px">
      <el-table :data="all_shop_account" height="500px" border="true">
        <el-table-column label="商店ID" prop="shopId" width="120"/>
        <el-table-column label="店主ID" prop="sellerUid" width="120"/>
        <el-table-column label="账户状态" prop="accountState" width="140"/>
        <el-table-column label="商品种类" prop="productType" width="140"/>
        <el-table-column label="商店余额" prop="balance" width="140"/>
      </el-table>
    </el-scrollbar>
  </div>
  <div>
    <el-card style="margin-top: 20px;width: 1100px">
      <el-form ref="checkForm" :model="recharge_shop" :rules="check_rules">
        <el-space :size="100" alignment="center" direction="horizontal">
          <el-form-item>
          <span style="font-size: medium">
            商店充值
          </span>
          </el-form-item>
          <el-form-item prop="shopId">
            <el-input
                v-model="recharge_shop.shopId"
                placeholder="请输入商店ID"
                prefix-icon="Edit"
                size="large"
            ></el-input>
          </el-form-item>
          <el-form-item prop="number">
            <el-input
                v-model="recharge_shop.number"
                placeholder="请输入充值金额"
                prefix-icon="Money"
                size="large"
            ></el-input>
          </el-form-item>
          <el-form-item style="text-align: right">
            <el-button autocomplete="off" type="primary" @click="rechargeShop"
            >确定
            </el-button
            >
          </el-form-item>
        </el-space>
      </el-form>
    </el-card>
  </div>
</template>

<script>
export default {
  name:"all_shop_account",
  data(){
    return{
      all_shop_account: localStorage.getItem('all_shop_account')
          ? JSON.parse(localStorage.getItem('all_shop_account'))
          : {},
      recharge_shop: {},
      check_rules: {
        shopId: [
          {required: true, message: '商店ID不能为空', trigger: 'blur'},
          {pattern: /^\d{5,}$/, message: '商店ID应为至少5位数字', trigger: 'blur'},
        ],
        number: [
          {required: true, message: '充值金额不能为空', trigger: 'blur'},
          {pattern: /^\d{1,}$/, message: '单次充值不小于10元', trigger: 'blur'},
        ],
      },
    }
  },
  created() {
    // 从后台获取最新的User数据
    this.findAll()
  },
  methods:{
    findAll(){
      this.request.get('/shop-account/find_all_shop_account').then((res) => {
        if (res.code === '200') {
          localStorage.setItem('all_shop_account', JSON.stringify(res.data)) // 存储用户信息到浏览器
        } else {
          this.$message.error(res.msg)
        }
      })
    },
    rechargeShop() {
      this.$refs['checkForm'].validate((valid) => {
        if (valid) {
          this.request.post('/shop-account/recharge_shop_account_'+this.recharge_shop.shopId, this.recharge_shop.number).then((res) => {
            if (res.code === '200') {
              this.$message.success("充值成功")
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