<template>
  <el-container style="height: 500px">
    <el-scrollbar height="500px">
      <el-table :data="all_product_data" border="true" height="500px">
        <el-table-column label="商品序列号" prop="productId" width="100"/>
        <el-table-column label="商品名" prop="productName" width="120"/>
        <el-table-column label="商品图片" prop="productImage" width="240">
          <template v-slot="scope">
            <div>
              <el-image v-for="(item, index) in scope.row.productImage"
                        :key="index" :src="item"
                        :preview-src-list="[item]"
                        alt="商品图片"
                        style="width: 50px;height: 50px"
                        append-to-body
                        :preview-teleported="true"
              >
              </el-image>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="商品价格" prop="productPrice" width="180"/>
        <el-table-column label="商品描述" prop="productInformation" width="120"/>
        <el-table-column label="商品状态" prop="productState" width="90"/>
        <el-table-column label="所属商店ID" prop="shopId" width="160"/>
      </el-table>
    </el-scrollbar>
  </el-container>

</template>

<script>

export default {
  name: "all_product_data",
  data() {
    return {
      all_product_data: localStorage.getItem('all_product_data')
          ? JSON.parse(localStorage.getItem('all_product_data'))
          : {},
    }
  },
  created() {
    this.findAll()
  },
  methods: {
    findAll() {
      this.request.get('/product-data/find_all_product_data').then((res) => {
        if (res.code === '200') {
          localStorage.setItem('all_product_data', JSON.stringify(res.data))
        } else {
          this.$message.error(res.msg)
        }
      })
    },
  }
}
</script>