import {
  type RouteConfig,
  layout,
  index,
  route,
  prefix,
} from '@react-router/dev/routes';

export default [
  layout('routes/layout.tsx', [
    index('routes/home.tsx'),
    route('/register', 'routes/register.tsx'),
    ...prefix('questionnaire', [
      route('/experience', 'routes/questionnaire/experience.tsx'),
      route('/personal-data', 'routes/questionnaire/personalData.tsx'),
    ]),
  ]),
] satisfies RouteConfig;
