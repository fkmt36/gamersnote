import { ActionContext } from 'vuex/types'
import { Context } from '@nuxt/types'
import { meStore } from '@/store'
import { $userApi } from '@/plugins/api'

export async function nuxtServerInit(
  _: ActionContext<any, any>,
  { res, redirect, route, req }: Context
) {
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
        res.setHeader('Set-Cookie', result.headers['set-cookie'])
      }
    } catch (err) {
      console.error(err)
    } finally {
      redirect('/')
    }
  } else if (route.path === '/verify-email') {
    try {
      if (
        typeof route.query.code === 'string' &&
        typeof route.query.uid === 'string'
      ) {
        await $userApi().patchUserEmailVerified({
          code: route.query.code,
          uid: route.query.uid,
        })
      }
    } catch (err) {
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
}
