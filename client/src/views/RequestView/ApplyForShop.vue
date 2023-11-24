<template>
  <div class="wrapper">
    <div
      style="
        margin: 100px auto;
        background-color: #fff;
        width: 350px;
        height: 480px;
        padding: 20px;
        border-radius: 10px;
      "
    >
      <div style="margin: 20px 0; text-align: center; font-size: 24px"><b>申请开店</b></div>
      <el-form :model="shop" :rules="rules" ref="shopForm">
        <el-form-item prop="shop_name">
          <el-input
            placeholder="请输入商店名"
            size="medium"
            prefix-icon="House"
            v-model="shop.shop_name"
          ></el-input>
        </el-form-item>
        <el-form-item prop="product_type">
          <el-input
            placeholder="请输入主营商品种类（以空格分隔多个种类）"
            size="medium"
            prefix-icon="Goods"
            v-model="shop.product_type"
          ></el-input>
        </el-form-item>
        <el-form-item prop="seller_id_number">
          <el-input
            placeholder="请输入店主身份证号"
            size="medium"
            prefix-icon="User"
            v-model="shop.seller_id_number"
          ></el-input>
        </el-form-item>
        <el-form-item prop="shop_resume">
          <el-input
            placeholder="请输入商店简历"
            size="medium"
            prefix-icon="EditPen"
            v-model="shop.shop_resume"
          ></el-input>
        </el-form-item>
        <el-form-item prop="shop_address">
          <el-input
            placeholder="请输入商店备案地址"
            size="medium"
            prefix-icon="Location"
            v-model="shop.shop_address"
          ></el-input>
        </el-form-item>
        <el-form-item prop="shop_register_cost">
          <el-input
            placeholder="请输入商店注册资金(单位：元)"
            size="medium"
            prefix-icon="Money"
            type="number"
            v-model="shop.shop_register_cost"
          ></el-input>
        </el-form-item>
        <el-form-item prop="shop_register_time">
          <el-date-picker
              v-model="shop.shop_register_time"
              type="datetime"
              placeholder="请选择商店注册时间"

          />
        </el-form-item>
        <el-form-item style="margin: 30px 0; text-align: right">
          <el-space direction="horizontal" alignment="center" :size="160">
            <el-button type="primary" size="medium" autocomplete="off" @click="apply"
              >提交申请</el-button
            >
            <el-button
              type="warning"
              size="medium"
              autocomplete="off"
              @click="$router.push('/seller')"
              >返回主界面</el-button
            >
          </el-space>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'apply_for_shop',
  data() {
    return {
      shop: {},
      rules: {
        shop_name: [
          { required: true, message: '商店名不能为空', trigger: 'blur' },
          { min: 1, max: 12, message: '长度应为 1 到 12 个字符', trigger: 'blur' }
        ],
        product_type: [{ required: true, message: '主营商品种类不能为空', trigger: 'blur' }],
        seller_id_number: [
          { required: true, message: '店主身份证号不能为空', trigger: 'blur' },
          { min: 18, max: 18, message: '长度应为 18 个字符', trigger: 'blur' },
          { pattern:/(^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)/, message: '身份证格式不正确', trigger: 'blur' }
        ],
        shop_resume: [
          { required: true, message: '商店简历不能为空', trigger: 'blur' },
          { min: 1, max: 120, message: '长度在 1 到 128 个字符', trigger: 'blur' }
        ],
        shop_address: [
          { required: true, message: '商店备案地址不能为空', trigger: 'blur' },
          { min: 1, max: 32, message: '长度在 1 到 32 个字符', trigger: 'blur' }
        ],
        shop_register_cost: [
            { required: true, message: '商店注册资金不能为空', trigger: 'blur' },
            { pattern: /^([1-9][0-9]{3,}.\d*)|([1-9][0-9]{3,})$/,message: '注册资金不低于1000元',trigger: 'blur'},
        ],
        shop_register_time: [{ required: true, message: '请选择注册时间', trigger: 'blur' }]
      },
    }
  },
  methods: {
    apply() {
      this.$refs['shopForm'].validate((valid) => {
        if (valid) {
          this.request.post('/shop-data/apply_for_shop', this.shop).then((res) => {
            if (res.code === '200') {
              this.$message.success('注册成功，审核中')
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

<style>
.wrapper {
  height: 100vh;
  background-color: #72767b;
  overflow: hidden;
}
</style>
