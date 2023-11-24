<template>
  <el-scrollbar height="605px">
    <el-table :data="shop_open_data" height="500px" border="true">
      <el-table-column label="商店名" prop="shopName" width="150"/>
      <el-table-column label="商品种类" prop="productType" width="250"/>
      <el-table-column label="简历" prop="shopResume" width="250"/>
      <el-table-column label="地址" prop="shopAddress" width="250"/>
      <el-table-column label="商店分数" prop="shopScore" width="100"/>
      <el-table-column label="商店状态" prop="shopState" width="100"/>
      <el-table-column label="店主UID" prop="sellerUid" />
    </el-table>
  </el-scrollbar>
</template>

<script>

export default {
  name:"shop_open_data",
  data(){
    return{
      shop_open_data: localStorage.getItem('shop_open_data')
          ? JSON.parse(localStorage.getItem('shop_open_data'))
          : {}
    }
  },
  created() {
    // 从后台获取最新的User数据
    this.findAll()
  },
  methods:{
    findAll(){
      this.request.get('/shop-data/find_open_shop').then((res) => {
        if (res.code === '200') {
          localStorage.setItem('shop_open_data', JSON.stringify(res.data)) // 存储用户信息到浏览器
        } else {
          this.$message.error(res.msg)
        }
      })
    }
  }
}
</script>