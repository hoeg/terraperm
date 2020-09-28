# terraperm

It is a best practice to use least priviledge policies when working with [AWS](https://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html#grant-least-privilege). However it can sometimes be hard to figure out which privileges that must be granted in order for terraform to be able to apply the infrastructure you have just created.

Often the workflow will be start with an empty policy:

- run terraform
- fail due to missing permission
- add permission to policy
- repeat

This process is tedious and timeconsuming. Terraperm aims to remove most of the burden in this task by giving you a policy with all the permissions needed for terraform to apply the infrastructure.

The new workflow using terraperm:

- grant all permissions to the services you are applying
- run terraperm from the same directory as you would run terraform
- update your policy with the output from terraperm


