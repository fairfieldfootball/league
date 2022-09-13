import { css } from '@emotion/react';

import { Theme } from './theme';

export default {
  body: function ({ palette }: Theme) {
    return css`
      body {
        min-height: 100vh;
        width: 100%;
        background-color: ${palette.document.background.base};
        color: ${palette.document.text.base};
        font-size: 100%;
        line-height: 1;
      }
    `;
  },
  borderBox: function () {
    return css`
    html {
      box-sizing: border-box;
    }
    * {
      box-sizing: inherit;
    }
    `;
  },
  fonts: function ({ typography }: Theme) {
    return css`
    body {
      html, body, input, select, optgroup, textarea, button {
        font-family: ${typography.fontAliases.text};
        font-size: ${typography.fontSizes.base};
        font-weight: ${typography.fontWeights.normal};
        font-feature-settings: 'kern';
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
        overflow-x: hidden;
        text-rendering: optimizeLegibility;
      }
    }
    `;
  },
  meyers: function () {
    return css`
    html, body, div, span, applet, object, iframe,
    h1, h2, h3, h4, h5, h6, p, blockquote, pre,
    a, abbr, acronym, address, big, cite, code,
    del, dfn, em, img, ins, kbd, q, s, samp,
    small, strike, strong, sub, sup, tt, var,
    b, u, i, center,
    dl, dt, dd, ol, ul, li,
    fieldset, form, label, legend,
    table, caption, tbody, tfoot, thead, tr, th, td,
    article, aside, canvas, details, embed,
    figure, figcaption, footer, header, hgroup,
    menu, nav, output, ruby, section, summary,
    time, mark, audio, video {
    	margin: 0;
    	padding: 0;
    	border: 0;
    	font-size: 100%;
    	font: inherit;
    	vertical-align: baseline;
    }
    /* HTML5 display-role reset for older browsers */
    article, aside, details, figcaption, figure,
    footer, header, hgroup, menu, nav, section {
    	display: block;
    }
    body {
    	line-height: 1;
    }
    ol, ul {
    	list-style: none;
    }
    blockquote, q {
    	quotes: none;
    }
    blockquote:before, blockquote:after,
    q:before, q:after {
    	content: '';
    	content: none;
    }
    table {
    	border-collapse: collapse;
    	border-spacing: 0;
    }
    `;
  },
};
