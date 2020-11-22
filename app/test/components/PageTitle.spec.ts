import { mount, Wrapper } from '@vue/test-utils'
// @ts-ignore
import PageTitle from '@/components/PageTitle.vue'

let wrapper: Wrapper<PageTitle>

afterEach(() => {
  wrapper.destroy()
})

describe('PageTitle', () => {
  it('mount properly', () => {
    wrapper = mount(PageTitle, { propsData: { title: '' } })
    expect(wrapper).toBeTruthy()
  })

  // スナップショットと比較
  // test('renders properly', () => {
  //   const wrapper = factory()
  //   expect(wrapper.html()).toMatchSnapshot()
  // })

  // タイトルの表示
  it('show title', () => {
    const title = 'テストタイトル'
    wrapper = mount(PageTitle, { propsData: { title } })
    expect(wrapper.find('h1').text()).toBe(title)
  })
})
