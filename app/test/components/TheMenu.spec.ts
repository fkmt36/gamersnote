import { mount, RouterLinkStub, Wrapper } from '@vue/test-utils'
// @ts-ignore
import TheMenu from '@/components/TheMenu.vue'
import { theMenuState } from '@/store'

let wrapper: Wrapper<TheMenu>

beforeEach(() => {
  wrapper = mount(TheMenu, {
    stubs: {
      NuxtLink: RouterLinkStub,
      'font-awesome-icon': true,
    },
  })
})

afterEach(() => {
  wrapper.destroy()
})

describe('TheMenu', () => {
  it('mount properly', () => {
    expect(wrapper).toBeTruthy()
  })

  // スナップショットと比較
  // test('renders properly', () => {
  //   const wrapper = factory()
  //   expect(wrapper.html()).toMatchSnapshot()
  // })

  // ホームに遷移
  describe('home', () => {
    it('transition home', () => {
      const e = wrapper.findAll('a').at(0)
      expect(e.text()).toBe('ホーム')
      expect(e.props().to).toBe('/')
    })
    it('close menu', () => {
      const e = wrapper.findAll('li').at(0)
      const spy = jest.spyOn(theMenuState, 'setShowMenu')
      e.trigger('click')
      expect(spy).toBeCalled()
    })
  })

  // ログイン・新規登録ページに遷移
  describe('login', () => {
    it('transition login', () => {
      const e = wrapper.findAll('a').at(1)
      expect(e.text()).toBe('ログイン・新規登録')
      expect(e.props().to).toBe('/login')
    })
    it('close menu', () => {
      const e = wrapper.findAll('li').at(1)
      const spy = jest.spyOn(theMenuState, 'setShowMenu')
      e.trigger('click')
      expect(spy).toBeCalled()
    })
  })

  // ボタン
  describe('Button', () => {
    it('close menu', () => {
      const e = wrapper.find('button')
      const spy = jest.spyOn(theMenuState, 'setShowMenu')
      e.trigger('click')
      expect(spy).toBeCalled()
    })
  })
})
