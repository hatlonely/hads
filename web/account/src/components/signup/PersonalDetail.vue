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
          <h2>欢迎使用 {{this.$config.org}}</h2>
        </v-flex>
        <v-flex my-4 xs12 lg12>
          <v-layout align-center justify-left>
            <v-card class="h-round-card px-2 py-1" outlined flat>
              <v-icon class="mr-2">account_circle</v-icon>
              <span class="body-1">{{this.$store.state.signup.email}}</span>
            </v-card>
          </v-layout>
        </v-flex>

        <v-flex mt-10 mb-4>
          <v-form ref="form" v-model="valid" lazy-validation>
            <v-layout mx-0 row wrap>
              <v-flex xs12>
                <v-menu
                  ref="birthdayMenu"
                  v-model="birthdayMenu"
                  :close-on-content-click="false"
                  transition="scale-transition"
                  offset-y
                  full-width
                  max-width="290px"
                  min-width="290px"
                >
                  <template v-slot:activator="{ on }">
                    <v-text-field
                      v-model="dateFormatted"
                      label="生日"
                      :rules="[rules.required]"
                      v-on="on"
                      outlined
                      filled
                      readonly
                      :value="dateFormatted = formatBirthday(birthday)"
                    ></v-text-field>
                  </template>
                  <v-date-picker
                    v-model="birthday"
                    no-title
                    ref="picker"
                    locale="zh-cn"
                    :max="new Date().toISOString().substr(0, 10)"
                    min="1950-01-01"
                    @change="save"
                    color="primary"
                    :day-format="date => date.split('-')[2]"
                  ></v-date-picker>
                </v-menu>
              </v-flex>
              <v-flex xs12>
                <v-select
                  v-model="gender"
                  :items="genderChoice"
                  item-text="val"
                  item-value="idx"
                  label="性别"
                  :rules="[rules.required]"
                  outlined
                  filled
                ></v-select>
              </v-flex>
            </v-layout>
          </v-form>
        </v-flex>
        <v-flex mt-0 mx-0>
          <v-layout align-left justify-center row fill-height text-left>
            <v-flex xs3>
              <v-btn text color="primary" @click="$router.go(-1)" pl-0>后退</v-btn>
            </v-flex>
            <v-flex xs6></v-flex>
            <v-flex xs3>
              <v-btn color="primary" depressed @click="validate" :disabled="!valid">下一步</v-btn>
            </v-flex>
          </v-layout>
        </v-flex>
      </v-flex>

      <v-flex xs4>
        <v-layout align-center justify-center row fill-height mx-0 class="text-center">
          <v-flex xs12>
            <v-flex xs12>
              <v-img :src="require('../../assets/img/signup_phone.svg')" max-width="244" inline></v-img>
            </v-flex>
            <v-flex xs12 px-4>
              <p class="body-1">您的个人信息不会对外公开，并且绝对安全</p>
            </v-flex>
          </v-flex>
        </v-layout>
      </v-flex>
    </v-layout>
  </v-card>
</template>

<script>
import rules from "../../assets/js/rules";

export default {
  methods: {
    validate() {
      if (this.$refs.form.validate()) {
        this.$router.push("/signup/privacy");
      }
    },
    formatBirthday(birthday) {
      if (!birthday) return null;
      const [year, month, day] = birthday.split("-");
      return `${year} 年 ${month} 月 ${day} 日`;
    },
    save(date) {
      this.$refs.birthdayMenu.save(date);
    }
  },
  computed: {
    birthday: {
      get() {
        return this.$store.state.signup.birthday;
      },
      set(birthday) {
        this.$store.commit("signup/setBirthday", birthday);
      }
    },
    gender: {
      get() {
        return this.$store.state.signup.gender;
      },
      set(gender) {
        this.$store.commit("signup/setGender", gender);
      }
    }
  },
  watch: {
    birthday(val) {
      this.dateFormatted = this.formatBirthday(val);
    },
    birthdayMenu(val) {
      val && setTimeout(() => (this.$refs.picker.activePicker = "YEAR"));
    }
  },
  data() {
    return {
      valid: true,
      dateFormatted: null,
      birthdayMenu: false,
      genderChoice: [
        { idx: 1, val: "男" },
        { idx: 2, val: "女" },
        { idx: 0, val: "保密" }
      ],
      rules
    };
  }
};
</script>
