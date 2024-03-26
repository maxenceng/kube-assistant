import kubeconfig from '@/app/reducers/kubeconfig'
import notification from '@/app/reducers/notification'
import { combineReducers } from 'redux'

export default combineReducers({
  kubeconfig,
  notification,
})
