import { mount, RouterLinkStub, Wrapper } from '@vue/test-utils'
// @ts-ignore
import TheHeader from '@/components/TheHeader.vue'
import { theMenuState } from '@/store'

let wrapper: Wrapper<TheHeader>

beforeEach(() => {
  wrapper = mount(TheHeader, {
    stubs: {
      NuxtLink: RouterLinkStub,
      'font-awesome-icon': true,
    },
  })
})

afterEach(() => {
  wrapper.destroy()
})

describe('TheHeader', () => {
  it('mount properly', () => {
    expect(wrapper).toBeTruthy()
  })

  // スナップショットと比較
  // test('renders properly', () => {
  //   const wrapper = factory()
  //   expect(wrapper.html()).toMatchSnapshot()
  // })

  // ロゴ
  describe('Logo', () => {
    it('push /', () => {
      expect(wrapper.find('a').text()).toBe('GamersNote')
      expect(wrapper.find('a').props().to).toBe('/')
    })
  })

  // ボタン
  describe('Button', () => {
    it('open menu', () => {
      wrapper.find('button').trigger('click')
      expect(theMenuState.getShowMenu).toBe(true)
    })
  })
})
