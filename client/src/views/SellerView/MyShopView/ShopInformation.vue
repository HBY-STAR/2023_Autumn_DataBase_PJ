<template>
  <div
      style="
        margin-top: 120px;
        margin-left: 350px;
        width: 800px;
        height: 520px;
        padding: 20px;
        border-radius: 10px;
        background-color: #909399;
      "
  >
    <div style="width: 400px; float: left">
      <el-card style="width: 420px;">
        <el-form label-width="90px" size="small">
          <div style="height: 30px">
            <header style="text-align: center; color: #337ecc">
              商店信息
            </header>
          </div>
          <el-form-item label="商店ID:">
            <span>{{ shop_data.shopId }}</span>
          </el-form-item>
          <el-form-item label="商店名:">
          <span>
            {{ shop_data.shopName }}
          </span>
          </el-form-item>
          <el-form-item label="商店简历:">
          <span>
            {{ shop_data.shopResume }}

          </span>
          </el-form-item>
          <el-form-item label="地址:">
          <span>
            {{ shop_data.shopAddress }}

          </span>
          </el-form-item>
          <el-form-item label="注册资金:">
          <span>
            {{ shop_data.shopRegisterCost }}
          </span>
          </el-form-item>
          <el-form-item label="注册时间:">
            <span>{{ shop_data.shopRegisterTime }}</span>
          </el-form-item>
          <el-form-item label="商店评分:">
            <span>{{ shop_data.shopScore }}</span>
          </el-form-item>
          <el-form-item label="店主UID:">
            <span>{{ shop_data.sellerUid }}</span>
          </el-form-item>
          <el-form-item label="店主身份证号:">
            <span>{{ shop_data.sellerIdNumber }}</span>
          </el-form-item>
          <el-form-item label="商店状态:">
            <span>{{ shop_data.shopState }}</span>
          </el-form-item>
          <el-form-item>
            <el-button
                autocomplete="off"
                style="position: absolute;right: 0; margin-top: 30px;margin-bottom: 20px"
                type="danger"
                @click="deleteShop"
            >注销商店
            </el-button
            >
          </el-form-item>
        </el-form>
      </el-card>
    </div>
    <div style="width: 100px; height: 100px; float: left">
    </div>
    <div style="width: 300px; float: left">
      <el-card style="width: 300px;">
        <el-form label-width="80px" size="small">
          <div style="height: 30px">
            <header style="text-align: center; color: #337ecc">
              商店账户
            </header>
          </div>
          <el-form-item label="账户ID:">
            <span>{{ shop_account.uid }}</span>
          </el-form-item>
          <el-form-item label="账户状态:">
            <span>{{ shop_account.accountState }}</span>
          </el-form-item>
          <el-form-item label="商品种类:">
            <span>{{ shop_account.productType }}</span>
          </el-form-item>
          <el-form-item label="账户余额:">
            <div style="width: 300px">
            <span>
              {{ shop_account.balance }}
              <RouterLink style="position: absolute;right: 0;" to="/seller_shop_information/recharge_shop_account">
                <el-icon size="medium">
                  <el-tooltip content="充值" effect="light" placement="right">
                    <money style="outline-width: 0">
                    </money>
                  </el-tooltip>
                </el-icon>
              </RouterLink>
            </span>
            </div>
            <RouterView name="shop_balance"></RouterView>
          </el-form-item>
        </el-form>
      </el-card>
      <el-form-item style="margin-top: 140px">
        <el-button
            autocomplete="off"
            style="position: absolute;right: 0; margin-top: 30px"
            type="warning"
            @click="returnToMain"
        >返回主界面
        </el-button
        >
      </el-form-item>
    </div>
  </div>
</template>

<script>
import {Money} from "@element-plus/icons-vue";

export default {
  name: "seller_shop_information",
  components: {Money},
  data() {
    return {
      user_data: localStorage.getItem('user_data')
          ? JSON.parse(localStorage.getItem('user_data'))
          : {},
      shop_data: localStorage.getItem('shop_data')
          ? JSON.parse(localStorage.getItem('shop_data'))
          : {
              shopId: 0,
              shopName: '',
              productType: '',
              shopResume: '',
              shopAddress: '',
              shopRegisterCost: 0.0,
              shopRegisterTime: '',
              shopScore: 0.0,
              sellerUid: 0,
              sellerIdNumber: '',
              shopState: ''
          },
      shop_account: localStorage.getItem('shop_account')
          ? JSON.parse(localStorage.getItem('shop_account'))
          : {
            shopId: 0,
            accountState: '',
            productType:'',
            productOrder:'',
            balance: 0.0,
          }
    }
  },
  created() {
    this.findData()
    this.findAccount()
  },
  methods: {
    findAccount() {
      this.request.get('/shop-account/find_by_uid_' + this.user_data.uid).then((res) => {
        if (res.code === '200') {
          localStorage.setItem('shop_account', JSON.stringify(res.data)) // 存储用户信息到浏览器
        } else {
          this.$message.error(res.msg)
        }
      })
    },
    findData(){
      this.request.get('/shop-data/find_by_uid_' + this.user_data.uid).then((res) => {
        if (res.code === '200') {
          localStorage.setItem('shop_data', JSON.stringify(res.data)) // 存储用户信息到浏览器
        } else {
          this.$message.error(res.msg)
        }
      })
    },
    returnToMain(){
      if (this.user_data.userType === 'sys_user') {
        this.$router.push('/sys_user')
      } else if (this.user_data.userType === 'user') {
        this.$router.push('/user')
      } else if (this.user_data.userType === 'seller') {
        this.$router.push('/seller')
      } else {
        this.$router.push('/')
      }
    },
    deleteShop(){
      this.$confirm("是否确认提交注销商店请求?", "警告", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      }).then(()=>{
        this.request.delete('/shop-data/delete_shop_'+this.shop_data.shopId).then(res=>{
          if(res.code==='200'){
            this.$message.success("注销请求已提交");
          }
          else{
            this.$message.error(res.msg);
          }
        })}
      ).catch(()=>{})
    },
  }
}
</script>

