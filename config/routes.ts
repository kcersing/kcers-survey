/**
 * @name umi 的路由配置
 * @description 只支持 path,component,routes,redirect,wrappers,name,icon 的配置
 * @param path  path 只支持两种占位符配置，第一种是动态参数 :id 的形式，第二种是 * 通配符，通配符只能出现路由字符串的最后。
 * @param component 配置 location 和 path 匹配后用于渲染的 React 组件路径。可以是绝对路径，也可以是相对路径，如果是相对路径，会从 src/pages 开始找起。
 * @param routes 配置子路由，通常在需要为多个路径增加 layout 组件时使用。
 * @param redirect 配置路由跳转
 * @param wrappers 配置路由组件的包装组件，通过包装组件可以为当前的路由组件组合进更多的功能。 比如，可以用于路由级别的权限校验
 * @param name 配置路由的标题，默认读取国际化文件 menu.ts 中 menu.xxxx 的值，如配置 name 为 login，则读取 menu.ts 中 menu.login 的取值作为标题
 * @param icon 配置路由的图标，取值参考 https://ant.design/components/icon-cn， 注意去除风格后缀和大小写，如想要配置图标为 <StepBackwardOutlined /> 则取值应为 stepBackward 或 StepBackward，如想要配置图标为 <UserOutlined /> 则取值应为 user 或者 User
 * @doc https://umijs.org/docs/guides/routes
 */
export default [
  {
    path: '/user',
    layout: false,
    routes: [
      {
        name: 'login',
        path: '/user/login',
        component: './user/login',
      },
    ],
  },
  // {
  //   path: '/welcome',
  //   name: 'welcome',
  //   icon: 'smile',
  //   component: './Welcome',
  // },
  // {
  //   path: '/admin',
  //   name: 'admin',
  //   icon: 'crown',
  //
  //   routes: [
  //     {
  //       path: '/admin',
  //       redirect: '/admin/sub-page',
  //     },
  //     {
  //       path: '/admin/sub-page',
  //       name: 'sub-page',
  //       component: './Admin',
  //     },
  //   ],
  // },
  // {
  //   name: 'list.table-list',
  //   icon: 'table',
  //   path: '/list',
  //   component: './TableList',
  // },
  {
    path: '/survey',
    name: '问卷管理',
    icon: 'form',
    routes: [
      { path: '/survey/list', component: '@/pages/survey/list', name: '问卷列表' },
      { path: '/survey/:id/design', component: '@/pages/survey/designer', name: '设计页面' , hideInMenu: true, },

      { path: '/survey/:id/respondent', component: '@/pages/survey/respondent', name: '统计分析' ,hideInMenu: true, },
    ],
  },
  { path: '/survey/:id/respondent', component: '@/pages/survey/respondent', name: '答题',layout: false,  },

  // {
  //   path: '/Survey',
  //   name: '问卷管理',
  //   icon: 'form',
  //   routes: [
  //     { path: '/Survey/list', component: '@/pages/Survey/List', name: '问卷列表' },
  //     { path: '/Survey/create', component: '@/pages/Survey/Edit', name: '创建问卷' },
  //     { path: '/Survey/:id/edit', component: '@/pages/Survey/Edit', name: '编辑问卷' },
  //     { path: '/Survey/:id/design', component: '@/pages/Survey/Designer', name: '设计问卷' },
  //     // { path: '/survey/:id/respond', component: '@/pages/survey/Respondent', name: '填写问卷' },
  //     // { path: '/survey/:id/statistics', component: '@/pages/survey/Statistics', name: '统计分析' },
  //   ],
  // },
  {
    path: '/',
    redirect: '/welcome',
  },
  {
    path: '*',
    layout: false,
    component: './404',
  },
];
