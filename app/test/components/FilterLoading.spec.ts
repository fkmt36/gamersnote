import { mount, Wrapper } from '@vue/test-utils'
// @ts-ignore
import FilterLoading from '@/components/FilterLoading.vue'
import { filterLoadingState } from '@/store'

let wrapper: Wrapper<FilterLoading>

afterEach(() => {
  wrapper.destroy()
})

describe('FilterLoading', () => {
  it('mount properly', () => {
    wrapper = mount(FilterLoading)
    expect(wrapper).toBeTruthy()
  })

  // スナップショットと比較
  // test('renders properly', () => {
  //   const wrapper = factory()
  //   expect(wrapper.html()).toMatchSnapshot()
  // })

  // 表示
  it('show', () => {
    const message = 'テストメッセージ'
    filterLoadingState.setLoading({
      isLoading: true,
      message,
    })
    wrapper = mount(FilterLoading)
    expect(wrapper.vm.show).toBe(true)
    expect(wrapper.find('.message').text()).toBe(message)
  })

  // 非表示
  it('do not show', () => {
    filterLoadingState.clearLoading()
    wrapper = mount(FilterLoading)
    wrapper = mount(FilterLoading)
    expect(wrapper.vm.show).toEqual(false)
  })
})
