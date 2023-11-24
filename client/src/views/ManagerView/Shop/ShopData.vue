<template>
  <el-scrollbar height="605px">
    <el-table :data="all_shop_data" height="605px" border="true">
      <el-button></el-button>
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
</template>

<script>

export default {
  name:"all_shop_data",
  data(){
    return{
      all_shop_data: localStorage.getItem('all_shop_data')
          ? JSON.parse(localStorage.getItem('all_shop_data'))
          : {}
    }
  },
  created() {
    // 从后台获取最新的User数据
    this.findAll()
  },
  methods:{
    findAll(){
      this.request.get('/shop-data/find_all_shop_data').then((res) => {
        if (res.code === '200') {
          localStorage.setItem('all_shop_data', JSON.stringify(res.data)) // 存储用户信息到浏览器
        } else {
          this.$message.error(res.msg)
        }
      })
    }
  }
}
</script>