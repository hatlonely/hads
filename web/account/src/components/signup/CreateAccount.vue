<template>
  <v-card width="748" height="511" flat outlined>
    <v-layout row wrap mx-0 text-left px-8>
      <v-flex xs8 pr-4>
        <v-flex mt-8 mb-5 xs12 lg12>
          <v-layout>
            <v-img :src="require('../../assets/logo.png')" max-width="70" inline></v-img>
          </v-layout>
        </v-flex>
        <v-flex my-4 xs12 lg12>
          <div class="headline">创建您的 hads 账号</div>
        </v-flex>
        <v-flex my-10>
          <v-layout mx-0 row wrap>
            <v-flex xs6 pr-4>
              <v-text-field label="姓氏" outlined filled></v-text-field>
            </v-flex>
            <v-flex xs6 pl-4>
              <v-text-field label="名字" outlined filled></v-text-field>
            </v-flex>
            <v-flex xs12>
              <v-text-field label="邮箱" outlined filled></v-text-field>
            </v-flex>
            <v-flex xs12>
              <v-text-field
                label="输入您的密码"
                :append-icon="show ? 'visibility' : 'visibility_off'"
                :type="show ? 'text' : 'password'"
                @click:append="show = !show"
                :rules="[rules.required, rules.min]"
                hint="使用8个或更多字符(字母、数字和符号的组合)"
                outlined
                filled
              ></v-text-field>
            </v-flex>
          </v-layout>
        </v-flex>
        <v-flex my-12 mx-12>
          <v-layout align-left justify-center row fill-height text-left>
            <v-flex xs3>
              <v-btn text color="primary" pl-0>登录现有账号</v-btn>
            </v-flex>
            <v-flex xs6></v-flex>
            <v-flex xs3>
              <v-btn color="primary" depressed to="/signin/password">下一步</v-btn>
            </v-flex>
          </v-layout>
        </v-flex>
      </v-flex>

      <v-flex xs4>
        <v-layout align-center justify-center row fill-height mx-0>
          <v-flex xs12>
            <v-flex xs12>
              <v-img :src="require('../../assets/create_account.svg')" max-width="244" inline></v-img>
            </v-flex>
            <v-flex xs12 px-4>
              <p class="body-1">只需要一个账号，您就可以使用 hads 的所有产品和服务</p>
            </v-flex>
          </v-flex>
        </v-layout>
      </v-flex>
    </v-layout>
  </v-card>
</template>

<script>
export default {
  data() {
    return {
      show: false,
      rules: {
        required: value => !!value || "Required.",
        min: v => v.length >= 8 || "至少8个字符"
      }
    };
  }
};
</script>
