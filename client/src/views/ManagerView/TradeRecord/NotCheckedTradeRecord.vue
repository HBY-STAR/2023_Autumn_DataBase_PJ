<template>
  <el-container style="height: 600px">
    <el-scrollbar height="600px">
      <el-table :data="not_checked_trade_record_data" border="true" height="600px">
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
  name: "not_checked_trade_record_data",
  data() {
    return {
      not_checked_trade_record_data: localStorage.getItem('not_checked_trade_record_data')
          ? JSON.parse(localStorage.getItem('not_checked_trade_record_data'))
          : {},
    }
  },
  created() {
    this.findAll()
  },
  methods: {
    findAll() {
      this.request.get('/all-trade-record/find_not_checked_order').then((res) => {
        if (res.code === '200') {
          localStorage.setItem('not_checked_trade_record_data', JSON.stringify(res.data))
        } else {
          this.$message.error(res.msg)
        }
      })
    },
  }
}
</script>