<template>
  <v-card flat max-width="900" color="#fafafa">
    <v-layout mt-5 text-center wrap>
      <v-flex mb-5 xs12 lg12>
        <div class="headline">个人信息</div>
      </v-flex>
      <v-flex mb-5 xs12 lg12>
        <div class="body-1">您在 {{this.$config.org}} 服务中使用的基本信息，例如您的姓名和照片</div>
      </v-flex>

      <v-flex xs0 lg1></v-flex>
      <v-flex mb-5 xs12 lg10>
        <!-- profile -->
        <v-card flat outlined class="px-4 py-4">
          <v-layout>
            <v-flex xs9>
              <v-card-title>个人资料</v-card-title>
              <v-card-text>
                <p class="text-lg-left">
                  使用 {{this.$config.org}} 服务的其他用户可能会看到部分信息。
                  <a href="#">了解详情</a>
                </p>
              </v-card-text>
            </v-flex>
            <v-flex xs3 mt-6>
              <v-avatar v-if="this.imageUrl" size="40" @click="pickFile">
                <v-img :src="this.imageUrl" max-width="40" max-height="40"></v-img>
              </v-avatar>
              <div v-else v-html="identicon" @click="pickFile"></div>

              <input
                type="file"
                style="display: none"
                ref="image"
                accept="image/*"
                @change="onFilePicked"
              />
            </v-flex>
          </v-layout>

          <v-list class="text-lg-left">
            <v-list-item-group>
              <v-list-item to="update/name">
                <v-flex xs3>
                  <div class="overline">姓名</div>
                </v-flex>
                <v-list-item-title>{{this.$store.state.account.lastName}} {{this.$store.state.account.firstName}}</v-list-item-title>
                <v-icon small>arrow_forward_ios</v-icon>
              </v-list-item>
              <v-divider class="ml-4"></v-divider>
              <v-list-item to="update/birthday">
                <v-flex xs3>
                  <div class="overline">生日</div>
                </v-flex>
                <v-list-item-title>{{formatBirthday(this.$store.state.account.birthday)}}</v-list-item-title>
                <v-icon small>arrow_forward_ios</v-icon>
              </v-list-item>
              <v-divider class="ml-4"></v-divider>
              <v-list-item to="update/gender">
                <v-flex xs3>
                  <div class="overline">性别</div>
                </v-flex>
                <v-list-item-title>{{formatGender(this.$store.state.account.gender)}}</v-list-item-title>
                <v-icon small>arrow_forward_ios</v-icon>
              </v-list-item>
              <v-divider class="ml-4"></v-divider>
              <v-list-item to="update/password">
                <v-flex xs3>
                  <div class="overline">密码</div>
                </v-flex>
                <v-list-item-title>********</v-list-item-title>
                <v-icon small>arrow_forward_ios</v-icon>
              </v-list-item>
            </v-list-item-group>
          </v-list>
        </v-card>

        <!-- contact info -->
        <v-card flat outlined class="mt-6 px-4 py-4">
          <v-card-title>联系方式</v-card-title>
          <v-list class="text-lg-left">
            <v-list-item-group>
              <v-list-item to="update/email">
                <v-flex xs3>
                  <div class="overline">邮箱</div>
                </v-flex>
                <v-list-item-title>{{this.$store.state.account.email}}</v-list-item-title>
                <v-icon small>arrow_forward_ios</v-icon>
              </v-list-item>
              <v-divider class="ml-4"></v-divider>
              <v-list-item to="update/phone">
                <v-flex xs3>
                  <div class="overline">电话</div>
                </v-flex>
                <v-list-item-title>{{this.$store.state.account.phone}}</v-list-item-title>
                <v-icon small>arrow_forward_ios</v-icon>
              </v-list-item>
            </v-list-item-group>
          </v-list>
        </v-card>
      </v-flex>
    </v-layout>
  </v-card>
</template>

<script>
const jdenticon = require("jdenticon");

export default {
  methods: {
    pickFile() {
      this.$refs.image.click();
    },
    onFilePicked(e) {
      const files = e.target.files;
      if (files[0] !== undefined) {
        this.imageName = files[0].name;
        console.log(this.imageName);
        if (this.imageName.lastIndexOf(".") <= 0) {
          return;
        }
        const fr = new FileReader();
        fr.readAsDataURL(files[0]);
        fr.addEventListener("load", () => {
          this.imageUrl = fr.result;
          this.imageFile = files[0]; // this is an image file that can be sent to server...
          console.log(this.imageUrl);
          console.log(this.imageFile);
        });
      } else {
        this.imageName = "";
        this.imageFile = "";
        this.imageUrl = "";
      }
      console.log(this.imageName);
      console.log(this.imageUrl);
      console.log(this.imageFile);
    },
    formatBirthday(birthday) {
      if (!birthday) return null;
      const [year, month, day] = birthday.split("-");
      return `${year} 年 ${month} 月 ${day} 日`;
    },
    formatGender(gender) {
      return {
        0: "保密",
        1: "男",
        2: "女"
      }[gender];
    }
  },
  computed: {
    identicon: function() {
      return jdenticon.toSvg(this.$store.state.account.email, 40);
    }
  },
  data: () => ({
    imageName: "",
    imageUrl: "",
    imageFile: ""
  })
};
</script>
