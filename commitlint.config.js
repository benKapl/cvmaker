// commitlint.config.js (at your repository root)
module.exports = {
	extends: ['@commitlint/config-conventional'],
	rules: {
		// Enforce specific types for your commits
		'type-enum': [
			2, // Error level (2 = error, 1 = warning, 0 = disabled)
			'always',
			[
				'feat', // New feature
				'fix', // Bug fix
				'docs', // Documentation changes
				'style', // Formatting, missing semi-colons, etc.; no code change
				'refactor', // A code change that neither fixes a bug nor adds a feature
				'test', // Adding missing tests or correcting existing tests
				'chore', // Other changes that don't modify src or test files
				'build', // Changes that affect the build system or external dependencies (e.g., go.mod, package.json)
				'ci', // Changes to our CI configuration files and scripts (e.g., GitHub Actions)
				'perf', // A code change that improves performance
				'revert' // Reverts a previous commit
			]
		],
		// Enforce specific scopes relevant to your project's codebases
		'scope-enum': [
			1, // Error level
			'always',
			[
				'api',      // Changes related to the Go backend
				'web',     // Changes related to the React/TypeScript frontend
				'db',           // Database schema or migration changes
				'ci',           // CI/CD configuration
				'docs',         // Project documentation (README, CONTRIBUTING, Notion links)
				'auth',         // Authentication related features
				'user-profile', // User profile management
				'offer-parsing',// Job offer parsing logic
				'resume-gen',   // Resume generation logic
				'deps',         // Dependency updates (e.g., Go modules, npm packages)
				'config',       // General configuration files
				'common'        // Shared utilities or types
			]
		],
		// Subject line rules
		'subject-empty': [2, 'never'], // Subject cannot be empty
		'subject-full-stop': [2, 'never', '.'], // Subject cannot end with a period
		'header-max-length': [2, 'always', 300] // Header max length (type(scope): subject)
	}
};
