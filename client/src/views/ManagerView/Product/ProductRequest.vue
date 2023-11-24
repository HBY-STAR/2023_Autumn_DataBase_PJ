<template>
  <el-container style="height: 500px">
    <el-scrollbar height="500px">
      <el-table :data="product_request_data" border="true" height="500px">
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
  <el-card style="margin-top: 30px">
    <el-form ref="checkForm" :model="check_product_request" :rules="check_rules">
      <el-space :size="100" alignment="center" direction="horizontal">
        <el-form-item>
        <span style="font-size: medium">
          处理上架商品请求
        </span>
        </el-form-item>
        <el-form-item prop="productId">
          <el-input
              v-model="check_product_request.productId"
              placeholder="请输入商品ID"
              prefix-icon="Edit"
              size="large"
          ></el-input>
        </el-form-item>
        <el-form-item prop="isPass">
          <el-radio-group v-model="check_product_request.isPass">
            <el-radio border label="通过">通过</el-radio>
            <el-radio border label="不通过">不通过</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item style="text-align: right">
          <el-button autocomplete="off" type="primary" @click="check_request"
          >确定
          </el-button
          >
        </el-form-item>
      </el-space>
    </el-form>
  </el-card>
</template>

<script>

export default {
  name: "shop_request_data",
  data() {
    return {
      check_product_request: {},
      check_rules: {
        productId: [
          {required: true, message: '商品序列号不能为空', trigger: 'blur'},
          {pattern: /^\d{6,}$/, message: '商品序列号应为至少6位数字', trigger: 'blur'},
        ],
        isPass: [
          {required: true, message: '请选择是否通过该请求', trigger: 'change'}
        ],
      },
      product_request_data: localStorage.getItem('product_request_data')
          ? JSON.parse(localStorage.getItem('product_request_data'))
          : {},
    }
  },
  created() {
    this.findAll()
  },
  methods: {
    findAll() {
      this.request.get('/product-data/find_request_product').then((res) => {
        if (res.code === '200') {
          localStorage.setItem('product_request_data', JSON.stringify(res.data))
        } else {
          this.$message.error(res.msg)
        }
      })
    },
    check_request() {
      this.$refs['checkForm'].validate((valid) => {
        if (valid) {
          this.request.post('/product-data/check_product_request', this.check_product_request).then((res) => {
            if (res.code === '200') {
              this.$message.success("该请求已处理")
              this.$router.post('/sys_user/product_request_data')
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