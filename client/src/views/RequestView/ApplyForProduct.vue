<template>
  <div class="wrapper">
    <div
        style="
        margin: 100px auto;
        background-color: #fff;
        width: 300px;
        height: 300px;
        padding: 20px;
        border-radius: 10px;
      "
    >
      <div style="margin: 20px 0; text-align: center; font-size: 24px"><b>上架商品</b></div>
      <el-form :model="product" :rules="rules" ref="productForm" style="margin-top: 30px">
        <el-form-item prop="productName">
          <el-input
              placeholder="请输入商品名"
              size="medium"
              prefix-icon="House"
              v-model="product.productName"
          ></el-input>
        </el-form-item>
        <el-form-item prop="productPrice">
          <el-input
              placeholder="请输入商品价格"
              size="medium"
              prefix-icon="Money"
              v-model="product.productPrice"
          ></el-input>
        </el-form-item>
        <el-form-item prop="productInformation">
          <el-input
              placeholder="请输入商品描述"
              size="medium"
              prefix-icon="EditPen"
              v-model="product.productInformation"
          ></el-input>
        </el-form-item>
        <el-form-item style="margin: 30px 0; text-align: right">
          <el-space direction="horizontal" alignment="center" :size="150">
            <el-button type="primary" size="medium" autocomplete="off" @click="apply"
            >提交申请</el-button
            >
            <el-button
                type="warning"
                size="medium"
                autocomplete="off"
                @click="$router.push('/seller/product_data')"
            >返回</el-button
            >
          </el-space>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'apply_for_product',
  data() {
    return {
      product: {

      },
      rules: {
        productName: [
          { required: true, message: '商品名不能为空', trigger: 'blur' },
          { min: 1, max: 12, message: '长度应为 1 到 12 个字符', trigger: 'blur' }
        ],
        productPrice: [
          { required: true, message: '商品价格不能为空', trigger: 'blur' },
          { pattern:/^([1-9][0-9]*)+(.[0-9]{1,2})?$/, message: '商品价格应为最多带两位小数的正数',trigger: 'blur'},
        ],
        productInformation: [
          { required: true, message: '商品描述不能为空', trigger: 'blur' },
          { min: 1, max: 128, message: '长度在 1 到 128 个字符', trigger: 'blur' }
        ],
      },
      shop_data: localStorage.getItem('shop_data')
          ? JSON.parse(localStorage.getItem('shop_data'))
          :{},
      user_data: localStorage.getItem('user_data')
          ? JSON.parse(localStorage.getItem('user_data'))
          : {},
      product_data:{
        productId:0,
        productName:'',
        productImage:'',
        productPrice:0.0,
        productInformation:'',
        productState:'',
        shopId:0,
      }
    }
  },
  created() {
    this.findShopId()
  },
  methods: {
    apply() {
      this.$refs['productForm'].validate((valid) => {
        if (valid) {
          this.request.post('/product-data/apply_for_product_by_'+this.shop_data.shopId,this.product).then((res) => {
            if (res.code === '200') {
              localStorage.setItem('product_data', JSON.stringify(res.data))
              this.$message.success('申请成功，审核中')
            } else {
              this.$message.error(res.msg)
            }
          })
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
    }
  }
}
</script>

<style>
.wrapper {
  height: 620px;
  background-color: #909399;
  overflow: hidden;
  border-radius: 10px;
}
</style>
