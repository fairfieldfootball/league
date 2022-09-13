import { createTheme } from '@mmdb/themes'

export const sizes = {
  components: {
    navbar: { height: '60px' },
  },
}

const theme = createTheme({
  palette: {
    contrastThreshold: 2.5,
    colors: {
      soi_blue: '#0074d9',
      soi_blue_dark: '#1c1073',
      soi_blue_light: '#88c32d',
      soi_green: '#88c32d',
      soi_orange: '#ffbe3d',
      soi_red: '#ff4136',
    },
    aliases: {
      primary: 'soi_blue_dark',
      secondary: 'soi_green',
      success: 'soi_green',
      error: 'soi_red',
      info: 'soi_blue',
      debug: 'blue_gray',
      warn: 'soi_orange',
    },
  },
  typography: {
    fonts: {
      comingSoon: 'Coming Soon',
      permanentMarker: 'Permanent Marker',
      gaegu: 'Gaegu',
    },
    fontAliases: {
      heading: 'comingSoon',
      subheading: 'permanentMarker',
      text: 'gaegu',
    },
  }
})

export type Theme = typeof theme;
export default theme;
