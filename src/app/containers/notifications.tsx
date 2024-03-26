import Notifications from '@/app/components/general/notifications'
import { getNotifications } from '@/app/selectors/notification'
import { useAppDispatch, useAppSelector } from '@/app/configs/store'
import { removeNotification } from '@/app/actions/notification'

export default () => {
  const notifications = useAppSelector(getNotifications)
  const dispatch = useAppDispatch()
  const remove = (uuid: string) => dispatch(removeNotification(uuid))
  return (
    <Notifications
      notifications={notifications}
      remove={remove}
    />
  )
}
