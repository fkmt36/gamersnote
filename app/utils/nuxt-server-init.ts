import { ActionContext } from 'vuex/types'
import { Context } from '@nuxt/types'
import { meStore } from '@/store'
import { $userApi } from '@/plugins/api'

export async function nuxtServerInit(
  _: ActionContext<any, any>,
  { res, redirect, route, req }: Context
) {
  // await $userApi
  //   .getMe({
  //     headers: {
  //       Cookie: 'session="Hello World!"',
  //     },
  //   })
  //   .catch((err) => {
  //     console.error(err)
  //   })
  // console.log(req.headers)
  // console.log(route.path)
  // res.setHeader('Set-Cookie', [`bbb=bbbbb; HttpOnly`])
  if (route.path === '/verify-code') {
    try {
      if (
        typeof route.query.code === 'string' &&
        typeof route.query.username === 'string'
      ) {
        const result = await $userApi().patchUserVerified({
          code: route.query.code,
          username: route.query.username,
        })
        console.log(result.headers)
        res.setHeader('Set-Cookie', result.headers['set-cookie'])
      }
    } catch (err) {
      console.error(err)
    } finally {
      redirect('/')
    }
  } else {
    try {
      const result = await $userApi().getMe({
        headers: {
          Cookie: req.headers.cookie,
        },
      })
      meStore.setMe(result.data)
    } catch (err) {}
  }
  // if (res && res.locals && res.locals.user) {
  //   authState.setToken(res.locals.user.idToken)
  //   await meStore.fetchMe()
  // }
}
