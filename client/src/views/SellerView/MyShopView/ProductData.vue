<template>
  <div style="width: 900px;float: left">
    <el-scrollbar height="605px">
      <el-table
          :data="my_shop_product_data"
          ref="product_table"
          height="605px"
          :border="true"
          highlight-current-row
      >
        <el-table-column label="商品序列号" prop="productId" width="100"/>
        <el-table-column label="商品名" prop="productName" width="120"/>
        <el-table-column label="商品图片" prop="productImage" width="210">
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
              <UploadFilled
                  style="height: 18px; width: 20px;position:absolute;right: 0; bottom:60%;text-align: center"
                  @click="this.focus_product_id=scope.row.productId;upload_product_image=true">
              </UploadFilled>
              <Delete
                  style="height: 18px; width: 20px;position:absolute;right: 0; bottom:10%;text-align: center"
                  @click="this.focus_product_id=scope.row.productId;delete_product_image=true">
              </Delete>
                <el-dialog
                    v-model="this.delete_product_image"
                    title="通知"
                    style="text-align: center"
                    width="20%"
                    :append-to-body="true"
                    :modal="false"
                >
                  <span>是否确定要删除该商品的商品图片？</span>
                  <template #footer>
                  <span class="dialog-footer">
                  <el-button @click="delete_product_image = false" style="position: absolute;left: 30px">取消</el-button>
                  <el-button type="primary" @click="deleteProductImage">
                    确定
                  </el-button>
                  </span>
                  </template>
                </el-dialog>
              <div>
            <el-dialog
                v-model="this.upload_product_image"
                title="上传商品图片"
                width="20%"
                style="text-align: center;"
                fit="fit"
                :append-to-body="true"
                :modal="false"
            >
              <div>
                <el-upload
                    ref="upload"
                    drag
                    multiple
                    list-type="picture"
                    enctype="multipart/form-data"
                    action="#"
                    :limit="1"
                    :show-file-list="true"
                    :auto-upload="false"
                    :http-request="picUpload"
                    :before-upload="beforeAvatarUpload"
                >

                  <div class="el-upload__text">
                    将文件拖至此处，或者 <em>点击上传</em>
                  </div>
                  <template #tip>
                    <div class="el-upload__tip">
                      文件格式为jpg或png，且不大于5MB。
                    </div>
                  </template>
                </el-upload>
              </div>
              <template #footer>
                    <span class="dialog-footer">
                      <el-button @click="this.upload_product_image = false" style="position: absolute;left: 30px">返回</el-button>
                      <el-button
                          type="primary"
                          @click="submitUpload();"
                      >
                        确认
                      </el-button>
                    </span>
              </template>
            </el-dialog>
              </div>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="价格" prop="productPrice" width="70"/>
        <el-table-column label="商品描述" prop="productInformation" width="280"/>
        <el-table-column label="商品状态" prop="productState" width="120"/>
      </el-table>
    </el-scrollbar>
  </div>
  <div style="float: left;margin-left: 50px">
    <el-card style="width: 300px;height: 300px">
      <div style="text-align: center;">
        <span style="color: #409eff">商品信息修改</span>
        <el-form :model="product" :rules="rules" ref="productForm" style="margin-top: 20px">
          <el-form-item prop="productId">
            <el-input
                placeholder="请输入商品序列号"
                size="default"
                prefix-icon="House"
                v-model="product.productId"
            ></el-input>
          </el-form-item>
          <el-form-item prop="productName">
            <el-input
                placeholder="请输入商品名"
                size="default"
                prefix-icon="House"
                v-model="product.productName"
            ></el-input>
          </el-form-item>
          <el-form-item prop="productPrice">
            <el-input
                placeholder="请输入商品价格"
                size="default"
                prefix-icon="Money"
                v-model="product.productPrice"
            ></el-input>
          </el-form-item>
          <el-form-item prop="productInformation">
            <el-input
                placeholder="请输入商品描述"
                size="default"
                prefix-icon="EditPen"
                v-model="product.productInformation"
            ></el-input>
          </el-form-item>
          <el-form-item style="margin: 20px 0; text-align: right">
              <el-button type="primary" size="default" autocomplete="off" @click="changeProduct" style="position: absolute;right: 0;margin-top: 20px"
              >提交修改</el-button
              >
          </el-form-item>
        </el-form>
      </div>
    </el-card>
    <el-card style="width: 300px;margin-top: 15px;height: 110px">
      <div>
        <div style="text-align: center">
        <span style="color: #409eff">上架新商品</span>
        </div>
        <el-button type="primary" size="default" autocomplete="off" @click="router().push('/seller/apply_for_product')" style="margin-top: 20px;position: absolute;right: 80px;"
        >申请</el-button
        >
      </div>
    </el-card>
    <el-card style="width: 300px;margin-top: 15px;height: 160px">
      <div>
        <div>
          <div style="text-align: center; height: 30px">
            <span style="color: #409eff">下架商品</span>
          </div>
          <div>
            <span style="width: 80px;float: left;font-size: small;margin-top: 15px;color: #c45656">商品序列号:</span>
            <el-input-number
                style="margin-top: 10px;width: 180px;"
                placeholder="请输入商品序列号"
                v-model="delete_product_id"
                controls-position="right"
            >
            </el-input-number>
          </div>
        </div>
        <el-button type="primary" size="default" autocomplete="off" @click="deleteProduct" style="margin-top: 20px; position: absolute;right: 80px;"
        >删除</el-button
        >
      </div>
    </el-card>
  </div>
</template>

<script>
import router from "@/router";
import {Delete, UploadFilled} from "@element-plus/icons-vue";

export default {
  name:"my_shop_product_data",
  components: {Delete, UploadFilled},
  data(){
    return{
      product: {},
      rules: {
        productId:[
          { required: true, message: '商品序列号不能为空', trigger: 'blur' },
          { pattern:/^[0-9]*$/ , message: '商品序列号应为数字',trigger: 'blur'},
        ],
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
      my_shop_product_data: localStorage.getItem('my_shop_product_data')
          ? JSON.parse(localStorage.getItem('my_shop_product_data'))
          : {},
      shop_data: localStorage.getItem('shop_data')
          ? JSON.parse(localStorage.getItem('shop_data'))
          :{},
      user_data: localStorage.getItem('user_data')
          ? JSON.parse(localStorage.getItem('user_data'))
          : {},
      delete_product_id:0,
      upload_product_image:false,
      delete_product_image:false,
      focus_product_id:0,
    }
  },
  created() {
    this.findShopId()
    this.findAll()
  },
  methods:{
    router() {
      return router
    },
    findAll(){
      this.request.get('/product-data/find_by_'+ this.shop_data.shopId).then((res) => {
        if (res.code === '200') {
          localStorage.setItem('my_shop_product_data', JSON.stringify(res.data)) // 存储用户信息到浏览器
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
    changeProduct(){
      this.$refs['productForm'].validate((valid) => {
        if (valid) {
          this.request.post('/product-data/change_product',this.product).then((res) => {
            if (res.code === '200') {
              localStorage.setItem('product_data', JSON.stringify(res.data))
              this.$message.success('修改已提交，申请中')
            } else {
              this.$message.error(res.msg)
            }
          })
        }
      })
    },
    deleteProduct(){
      this.request.delete('/product-data/delete_'+this.delete_product_id).then((res) => {
        if (res === true) {
          this.$message.success('删除成功')
        } else {
          this.$message.error("删除失败")
        }
      })
    },
    beforeAvatarUpload(file) {
      const isJPG = file.type === 'image/jpeg';
      const isPNG = file.type === 'image/png';
      const isLt5M = file.size / 1024 / 1024 < 5;

      if ((!isPNG)&&(!isJPG)) {
        this.$message.error('上传头像图片只能是 JPG 或 PNG 格式!');
      }
      if (!isLt5M) {
        this.$message.error('上传头像图片大小不能超过 5MB!');
      }
      return (isJPG || isPNG) && isLt5M;
    },
//这里是自定义发送请求
    picUpload(f) {
      let params = new FormData()
      //注意在这里一个坑f.file
      params.append("file", f.file);
      this.$axios({
        method:'post',
        //这里的id是我要改变用户的ID值
        url:'/product-data/upload_image_'+this.focus_product_id,
        data:params,
        headers:{
          'content-type':'multipart/form-data'
        }
      }).then(res => {
          if (res.data.code==='200') {
            this.$message.success("上传成功");
            this.my_shop_product_data.productImage=res.data.data;
          }else {
            this.$message.error(res.data.msg);
          }
      })
    },
    //触发请求
    submitUpload() {
      this.$refs.upload.submit();
    },
    deleteProductImage(){
      this.request.delete('/product-data/delete_image_'+this.focus_product_id).then(res=>{
        if (res === true) {
          this.$message.success('删除成功');
        } else {
          this.$message.error("删除失败");
        }
      })
    },
  }
}
</script>