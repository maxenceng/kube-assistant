import { NotificationMessage } from '@/app/models/notification'
import { Alert, Fade, Grid } from '@mui/material'

interface Props {
  notifications: Array<NotificationMessage>
  remove: (uuid: string) => any
}

export default ({ notifications, remove }: Props) => (
  <Grid container spacing={2}>
    {notifications.map((notification) => (
      <Fade in>
        <Grid item key={notification.uuid}>
          <Alert
            onClose={() => remove(notification.uuid)}
            severity={notification.type}
          >
            {notification.message}
          </Alert>
        </Grid>
      </Fade>
    ))}
  </Grid>
)
