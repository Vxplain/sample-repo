# VX fixture scenarios

These pull requests are stable test data. Preserve their documented state.

| PR | Scenario | Expected state |
| --- | --- | --- |
| [#1](https://github.com/Vxplain/sample-repo/pull/1) | Merged documentation | Merged with a passing workflow and merge timeline |
| [#2](https://github.com/Vxplain/sample-repo/pull/2) | Abandoned experiment | Closed without merging, with an issue comment |
| [#3](https://github.com/Vxplain/sample-repo/pull/3) | Draft pagination | Open draft with passing and skipped checks |
| [#4](https://github.com/Vxplain/sample-repo/pull/4) | Mixed check results | Open with pass, fail, and skipped conclusions |
| [#5](https://github.com/Vxplain/sample-repo/pull/5) | Rich review | Open with multiple commits, Markdown, Unicode, requested reviewer, milestone, and current, resolved, outdated, replied-to, and pending threads |
| [#6](https://github.com/Vxplain/sample-repo/pull/6) | Unusual files | Open with a 25,000-row generated diff, binary image, symlink, executable, and missing final newline |
| [#7](https://github.com/Vxplain/sample-repo/pull/7) | Merge conflict | Open and conflicting with `main` |
| [#8](https://github.com/Vxplain/sample-repo/pull/8) | Pure rename | Open with one 100% file rename |
| [#9](https://github.com/Vxplain/sample-repo/pull/9) | Pending approval | Open with an environment-gated check waiting indefinitely |
| [#10](https://github.com/Vxplain/sample-repo/pull/10) | No checks | Open with an empty check rollup |

## Checks

The `Sample checks` workflow provides:

- `Unit tests` and `Go vet` as normal passing checks.
- `Fixture policy`, controlled by `.github/check-scenario`, as a passing or failing check.
- `Optional slow check`, normally skipped. Add the `check:slow` label to observe a five-minute in-progress check.
- `Approval gate`, enabled by `check:pending` and held by the `fixture-approval` environment without consuming runner time.

## Maintenance

- Keep `main` green.
- Do not merge or close open scenario pull requests.
- Do not resolve #7.
- Do not approve the deployment on #9.
- Do not submit the pending review on #5.
- Add new scenarios instead of repurposing existing ones.
