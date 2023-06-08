module.exports = {
  branches: [
    { name: "main" }, 
    { name: "next" },
    { name: "ci/improve-no-pr", prerelease: "2.29-ignore-me" }
  ],
  plugins: [
    "@semantic-release/commit-analyzer",
    "@semantic-release/release-notes-generator",
    "@semantic-release/github"
  ],
};
