import { KubeconfigLocation } from '@/app/models/kubeconfig'
import {
  Avatar, ListItemButton,
  ListItemIcon,
  ListItemText,
  MenuItem,
  MenuList,
} from '@mui/material'
import { green, grey } from '@mui/material/colors'
import { MoreVert } from '@mui/icons-material'

interface Props {
  kubeconfigLocations: Array<KubeconfigLocation>
  connect: (location: string) => void
}

// FIXME: Set context menu
export default ({ kubeconfigLocations, connect } : Props) => (
  <MenuList>
    {kubeconfigLocations.map((kubeconfigLocation) => (
      <MenuItem onClick={() => connect(kubeconfigLocation.location)}>
        <ListItemIcon>
          <Avatar
            sx={{ bgcolor: kubeconfigLocation.connected ? green : grey }}
          >
            {kubeconfigLocation.name.at(0)}
          </Avatar>
        </ListItemIcon>
        <ListItemText>
          {kubeconfigLocation.name}
        </ListItemText>
        <ListItemButton>
          <MoreVert />
        </ListItemButton>
      </MenuItem>
    ))}
  </MenuList>
)
