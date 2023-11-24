<template>
  <el-container style="height: 500px">
    <el-scrollbar height="500px">
      <el-table :data="all_sell_product_data" border="true" height="500px">
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
        <el-table-column label="添加到购物车" width="120">
          <template v-slot="scope">
            <div style="text-align: center">
              <ShoppingCart
                  style="height: 25px; width: 25px;text-align: center;color: #409eff"
                  @click="this.focus_product_id=scope.row.productId;addToShoppingCart();">
              </ShoppingCart>
            </div>
          </template>
        </el-table-column>
      </el-table>
    </el-scrollbar>
  </el-container>
</template>

<script>

import {ShoppingCart} from "@element-plus/icons-vue";

export default {
  name: "all_sell_product_data",
  components: {ShoppingCart},
  data() {
    return {
      all_sell_product_data: localStorage.getItem('all_sell_product_data')
          ? JSON.parse(localStorage.getItem('all_sell_product_data'))
          : {},
      user_data: localStorage.getItem('user_data')
          ? JSON.parse(localStorage.getItem('user_data'))
          : {},
      user_account: localStorage.getItem('user_account')
          ? JSON.parse(localStorage.getItem('user_account'))
          : {},
      focus_product_id:0,
    }
  },
  created() {
    this.findAll()
  },
  methods: {
    findAll() {
      this.request.get('/product-data/find_all_sell_product_data').then((res) => {
        if (res.code === '200') {
          localStorage.setItem('all_sell_product_data', JSON.stringify(res.data))
        } else {
          this.$message.error(res.msg)
        }
      })
    },
    addToShoppingCart(){
      if(this.user_data.uid==null){
        this.$message.error("请先登录");
      }else{
        this.request.post('/user-account/add_to_shopping_cart_'+this.user_data.uid,this.focus_product_id).then(res=>{
          if(res.code==='200'){
            localStorage.setItem('user_account',JSON.stringify(res.data))
            this.$message.success("添加成功！")
          }
          else {
            this.$message.error(res.msg)
          }
        })
      }

    }
  }
}
</script>