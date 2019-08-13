<template>
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
      <v-card flat outlined>
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
            <v-avatar size="50">
              <img :src="require('../../assets/img/hatlonely.png')" alt="avatar" />
            </v-avatar>
          </v-flex>
        </v-layout>

        <v-list class="text-lg-left">
          <v-list-item-group>
            <v-divider class="ml-4"></v-divider>
            <v-list-item>
              <v-flex xs3>
                <div class="overline">姓名</div>
              </v-flex>
              <v-list-item-title>{{this.$store.state.account.lastName}} {{this.$store.state.account.firstName}}</v-list-item-title>
              <v-icon small>arrow_forward_ios</v-icon>
            </v-list-item>
            <v-divider class="ml-4"></v-divider>
            <v-list-item>
              <v-flex xs3>
                <div class="overline">生日</div>
              </v-flex>
              <v-list-item-title>May 15, 1992</v-list-item-title>
              <v-icon small>arrow_forward_ios</v-icon>
            </v-list-item>
            <v-divider class="ml-4"></v-divider>
            <v-list-item>
              <v-flex xs3>
                <div class="overline">性别</div>
              </v-flex>
              <v-list-item-title>Male</v-list-item-title>
              <v-icon small>arrow_forward_ios</v-icon>
            </v-list-item>
            <v-divider class="ml-4"></v-divider>
            <v-list-item>
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
      <v-card flat outlined class="mt-6">
        <v-card-title>联系方式</v-card-title>

        <v-list class="text-lg-left">
          <v-list-item-group>
            <v-list-item>
              <v-flex xs3>
                <div class="overline">邮箱</div>
              </v-flex>
              <v-list-item-title>{{this.$store.state.account.email}}</v-list-item-title>
              <v-icon small>arrow_forward_ios</v-icon>
            </v-list-item>
            <v-divider class="ml-4"></v-divider>
            <v-list-item>
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
</template>

<script>
export default {
  data: () => ({})
};
</script>
