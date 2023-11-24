<template>
  <div>
    <el-scrollbar height="500px">
      <el-table :data="all_user_account" height="500px" border="true">
        <el-table-column label="用户序列号" prop="uid" width="120"/>
        <el-table-column label="账户状态" prop="accountState" width="120"/>
        <el-table-column label="账户余额" prop="balance" width="140"/>
      </el-table>
    </el-scrollbar>
  </div>
  <div>
    <el-card style="margin-top: 20px;width: 1100px">
      <el-form ref="checkForm" :model="recharge_user" :rules="check_rules">
        <el-space :size="100" alignment="center" direction="horizontal">
          <el-form-item>
          <span style="font-size: medium">
            个人账户充值
          </span>
          </el-form-item>
          <el-form-item prop="uid">
            <el-input
                v-model="recharge_user.uid"
                placeholder="请输入用户序列号"
                prefix-icon="Edit"
                size="large"
            ></el-input>
          </el-form-item>
          <el-form-item prop="number">
            <el-input
                v-model="recharge_user.number"
                placeholder="请输入充值金额"
                prefix-icon="Money"
                size="large"
            ></el-input>
          </el-form-item>
          <el-form-item style="text-align: right">
            <el-button autocomplete="off" type="primary" @click="rechargeUser"
            >确定
            </el-button
            >
          </el-form-item>
        </el-space>
      </el-form>
    </el-card>
  </div>
</template>

<script>
export default {
  name:"all_user_account",
  data(){
    return{
      all_user_account: localStorage.getItem('all_user_account')
          ? JSON.parse(localStorage.getItem('all_user_account'))
          : {},
      recharge_user: {},
      check_rules: {
        uid: [
          {required: true, message: '用户序列号不能为空', trigger: 'blur'},
          {pattern: /^\d{4,}$/, message: '用户序列号应为至少4位数字', trigger: 'blur'},
        ],
        number: [
          {required: true, message: '充值金额不能为空', trigger: 'blur'},
          {pattern: /^\d{1,}$/, message: '单次充值不小于10元', trigger: 'blur'},
        ],
      },
    }
  },
  created() {
    // 从后台获取最新的User数据
    this.findAll()
  },
  methods:{
    findAll(){
      this.request.get('/user-account/find_all_user_account').then((res) => {
        if (res.code === '200') {
          localStorage.setItem('all_user_account', JSON.stringify(res.data)) // 存储用户信息到浏览器
        } else {
          this.$message.error(res.msg)
        }
      })
    },
    rechargeUser() {
      this.$refs['checkForm'].validate((valid) => {
        if (valid) {
          this.request.post('/user-account/recharge_account', this.recharge_user).then((res) => {
            if (res.code === '200') {
              this.$message.success("充值成功")
            } else {
              this.$message.error(res.msg)
            }
          })
        }
      })
    },
  }
}
</script>