<template>
  <el-container style="height: 600px">
    <el-scrollbar height="600px">
      <el-table :data="trade_record_data" border="true" height="600px">
        <el-table-column label="订单序列号" prop="recordId" width="100"/>
        <el-table-column label="用户ID" prop="uid" width="200"/>
        <el-table-column label="商品ID" prop="productId" width="200"/>
        <el-table-column label="商品价格" prop="price" width="180"/>
        <el-table-column label="订单状态" prop="state" width="240"/>
        <el-table-column label="创建时间" prop="time" width="200"/>
      </el-table>
    </el-scrollbar>
  </el-container>
</template>

<script>

export default {
  name: "trade_record_data",
  data() {
    return {
      trade_record_data: localStorage.getItem('trade_record_data')
          ? JSON.parse(localStorage.getItem('trade_record_data'))
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
      this.request.get('/all-trade-record/find_checked_order_by_'+this.shop_data.shopId).then((res) => {
        if (res.code === '200') {
          localStorage.setItem('trade_record_data', JSON.stringify(res.data))
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
  }
}
</script>