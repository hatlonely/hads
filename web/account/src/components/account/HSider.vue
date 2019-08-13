<template>
  <v-navigation-drawer floating color="rgb(0, 0, 0, 0)" app cliped>
    <v-flex xs12>
      <v-card height="100" color="rgb(0, 0, 0, 0)" flat></v-card>
    </v-flex>
    <v-list dense rounded>
      <v-list-item-group v-model="idx" color="primary">
        <template v-for="(item, i) in items">
          <template v-if="item.divider">
            <v-divider :key="i"></v-divider>
          </template>
          <template v-else>
            <v-list-item :to="item.to" :key="i">
              <v-list-item-icon>
                <v-icon filled>{{ item.icon }}</v-icon>
              </v-list-item-icon>
              <v-list-item-content>
                <v-list-item-title>
                  <span>{{ item.title }}</span>
                </v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </template>
        </template>
      </v-list-item-group>
    </v-list>

    <v-footer absolute color="rgba(0, 0, 0, 0)">
      <v-flex text-center xs12>
        <div class="caption">
          <a href="#">隐私</a>&nbsp;&nbsp;&nbsp;&nbsp;
          <a href="#">条款</a>&nbsp;&nbsp;&nbsp;&nbsp;
          <a href="#">帮助</a>
        </div>
      </v-flex>
    </v-footer>
  </v-navigation-drawer>
</template>

<script>
export default {
  data() {
    return {
      idx: 1,
      items: [
        { title: "首页", icon: "account_circle", to: "/account/home" },
        {
          title: "个人信息",
          icon: "assignment_ind",
          to: "/account/personinfo"
        },
        { title: "数据和个性化", icon: "toggle_on", to: "/" },
        { title: "安全性", icon: "lock", to: "/" },
        { title: "用户和分享", icon: "people", to: "/" },
        { title: "付费和订阅", icon: "payment", to: "/" },
        { divider: true },
        { title: "帮助", icon: "help_outline", to: "/" },
        { title: "发送反馈", icon: "feedback", to: "/" }
      ]
    };
  }
};
</script>
