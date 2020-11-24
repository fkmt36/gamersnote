import { AxiosError } from 'axios'

export function isAxiosError<T>(
  error: AxiosError | any
): error is AxiosError<T> {
  return error && error.isAxiosError
}
