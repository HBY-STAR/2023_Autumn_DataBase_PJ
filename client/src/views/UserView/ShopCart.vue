<template>
  <el-container style="height: 500px">
    <el-scrollbar height="500px">
      <el-table :data="shop_cart_product"
                stripe
                @selection-change="handleSelectionChange"
                border="true"
                height="500px"
      >

        <el-table-column type="selection" width="50"/>
        <el-table-column label="商品序列号" prop="productId" width="100"/>
        <el-table-column label="商品名" prop="productName" width="120"/>
        <el-table-column label="商品图片" prop="productImage" width="280">
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
  <el-container>
    <div>
      <el-card style="margin-top: 20px;height: 80px;width: 1100px">
        <div>
          <div style="float: left">
            <el-button
                size="default"
                @click="handleDelete()"
                type="danger"
                style="margin-left:10px;margin-top: 3px;"
                icon="Delete"
                :disabled="multiple"
            >删除</el-button>
          </div>
          <div style="height: 50px; width: 800px;float: left">
          </div>
          <div>
            <el-button
                size="default"
                @click="handlePay()"
                type="success"
                style="margin-left:10px;margin-top: 3px;"
                icon="CircleCheck"
                :disabled="multiple"
            >结算</el-button>
          </div>

        </div>
      </el-card>
    </div>
  </el-container>

</template>

<script>
import QS from 'qs'

export default {
  name: "shop_cart",
  data() {
    return {
      shop_cart_product: localStorage.getItem('shop_cart_product')
          ? JSON.parse(localStorage.getItem('shop_cart_product'))
          : {},
      user_data: localStorage.getItem('user_data')
          ? JSON.parse(localStorage.getItem('user_data'))
          : {},
      // 选中数组
      productIds: [],
      // 非单个禁用
      single: true,
      // 非多个禁用
      multiple: true,
    }
  },
  created() {
    this.findByUid()
  },
  methods: {
    findByUid() {
      this.request.get('/product-data/find_shop_cart_by_'+this.user_data.uid).then((res) => {
        if (res.code === '200') {
          localStorage.setItem('shop_cart_product', JSON.stringify(res.data))
        } else {
          this.$message.error(res.msg)
        }
      })
    },
    handleSelectionChange(selection) {
      console.log("多选框选中数据");
      this.productIds = selection.map(item => item.productId);
      this.single = selection.length !== 1;
      this.multiple = !selection.length;
    },
    handleDelete(){
      this.$confirm("是否确认删除选中的商品?", "警告", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      }).then(()=>{
          this.$axios({
            method: "post",
            url: '/user-account/delete_shop_cart_by_'+this.user_data.uid,
            data: QS.stringify({productIds:this.productIds}, {arrayFormat: 'brackets'})
          }).then(res=>{
            if(res.data.code==='200'){
              this.$message.success("删除成功");
            }
            else{
              this.$message.error(res.data.msg);
            }
          })}
      ).catch(()=>{})

    },
    handlePay(){
      this.$confirm("是否结算选中的商品?", "通知", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      }).then(()=>{
        this.$axios({
          method: "post",
          url: '/user-account/pay_shop_cart_by_'+this.user_data.uid,
          data: QS.stringify({productIds:this.productIds}, {arrayFormat: 'brackets'})
        }).then(res=>{
          if(res.data.code==='200'){
            this.$message.success("购买成功");
          }
          else{
            this.$message.error(res.data.msg);
          }
        })}
      ).catch(()=>{})
    },
  }
}
</script>