<template>
  <v-card width="748" height="511" flat outlined>
    <v-layout row wrap mx-0 text-left px-8>
      <v-flex xs8 pr-8>
        <v-flex mt-8 mb-5 xs12 lg12>
          <v-layout>
            <v-img :src="require('../../assets/img/logo.png')" max-width="70" inline></v-img>
          </v-layout>
        </v-flex>
        <v-flex my-4 xs12 lg12>
          <h2>创建您的 {{this.$config.org}} 账号</h2>
        </v-flex>
        <v-flex mt-10 mb-4>
          <v-form ref="form" v-model="valid" lazy-validation>
            <v-layout mx-0 row wrap>
              <v-flex xs6 pr-4>
                <v-text-field
                  v-model="lastName"
                  label="姓氏"
                  :rules="[rules.required]"
                  outlined
                  filled
                  validate-on-blur
                ></v-text-field>
              </v-flex>
              <v-flex xs6 pl-4>
                <v-text-field
                  v-model="firstName"
                  label="名字"
                  :rules="[rules.required]"
                  outlined
                  filled
                  validate-on-blur
                ></v-text-field>
              </v-flex>
              <v-flex xs12>
                <v-text-field
                  v-model="email"
                  label="邮箱"
                  :rules="[rules.required, rules.validemail]"
                  :error-messages="errors"
                  outlined
                  filled
                  validate-on-blur
                ></v-text-field>
              </v-flex>
              <v-flex xs12>
                <v-text-field
                  v-model="password"
                  label="输入您的密码"
                  :append-icon="show ? 'visibility' : 'visibility_off'"
                  :type="show ? 'text' : 'password'"
                  @click:append="show = !show"
                  :rules="[rules.required, rules.atleast8characters]"
                  hint="使用8个或更多字符(字母、数字和符号的组合)"
                  outlined
                  filled
                  validate-on-blur
                ></v-text-field>
              </v-flex>
            </v-layout>
          </v-form>
        </v-flex>
        <v-flex mt-0 mx-0>
          <v-layout align-left justify-center row fill-height text-left>
            <v-flex xs3>
              <v-btn text color="primary" pl-0 to="/signin">登录现有账号</v-btn>
            </v-flex>
            <v-flex xs6></v-flex>
            <v-flex xs3>
              <v-btn :disabled="!valid" color="primary" depressed @click="validate">下一步</v-btn>
            </v-flex>
          </v-layout>
        </v-flex>
      </v-flex>

      <v-flex xs4>
        <v-layout align-center justify-center row fill-height mx-0 class="text-center">
          <v-flex xs12>
            <v-flex xs12>
              <v-img :src="require('../../assets/img/create_account.svg')" max-width="244" inline></v-img>
            </v-flex>
            <v-flex xs12 px-4>
              <p class="body-1">只需要一个账号，您就可以使用 {{this.$config.org}} 的所有产品和服务</p>
            </v-flex>
          </v-flex>
        </v-layout>
      </v-flex>
    </v-layout>
  </v-card>
</template>

<script>
const axios = require("axios");
import rules from "../../assets/js/rules";

export default {
  methods: {
    validate() {
      if (this.$refs.form.validate()) {
        this.$router.push("/signup/verifyphone");
      }
    }
  },
  watch: {
    email(val) {
      if (/^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}$/.test(val)) {
        axios
          .get(this.$config.api + "/vertify", {
            params: {
              field: "email",
              value: val
            },
            withCredentials: true
          })
          .then(res => {
            this.errors = res.data.ok ? [] : [res.data.tip];
          })
          .catch(function(error) {
            this.errors = error;
          });
      }
    }
  },
  computed: {
    email: {
      get() {
        return this.$store.state.signup.email;
      },
      set(email) {
        this.$store.state.signup.email = email;
      }
    },
    firstName: {
      get() {
        return this.$store.state.signup.firstName;
      },
      set(firstName) {
        this.$store.state.signup.firstName = firstName;
      }
    },
    lastName: {
      get() {
        return this.$store.state.signup.lastName;
      },
      set(lastName) {
        this.$store.state.signup.lastName = lastName;
      }
    },
    password: {
      get() {
        return this.$store.state.signup.password;
      },
      set(password) {
        this.$store.state.signup.password = password;
      }
    }
  },
  data() {
    return {
      valid: true,
      show: false,
      errors: [],
      rules
    };
  }
};
</script>
