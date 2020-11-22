import { mount, Wrapper } from '@vue/test-utils'
// @ts-ignore
import FilterDarken from '@/components/FilterDarken.vue'
import { filterDarkenState } from '@/store'

describe('FilterDarken', () => {
  it('mount properly', () => {
    const wrapper = mount(FilterDarken)
    expect(wrapper).toBeTruthy()
  })

  // スナップショットと比較
  // test('renders properly', () => {
  //   const wrapper = factory()
  //   expect(wrapper.html()).toMatchSnapshot()
  // })

  // 表示
  it('show', () => {
    filterDarkenState.setDarken(true)
    const wrapper: Wrapper<FilterDarken> = mount(FilterDarken)
    expect(wrapper.vm.show).toBe(true)
  })

  // 非表示
  it('do not show', () => {
    filterDarkenState.setDarken(false)
    const wrapper: Wrapper<FilterDarken> = mount(FilterDarken)
    expect(wrapper.vm.show).toBe(false)
  })
})
