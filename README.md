
## GB-Emulator

An emulator written for the DMG-Gameboy in golang. I wrote this to get more familiar with writing larger projects.
As of now Tetris is the only game that runs. MBC1 support has partially been implemented

<img src="./images/tetris.png" width="800" height="450" alt="tetris_screenshit"/> \

**Blargg's Rom tests**
| Test          | Status  |
| --------      | ------- |
| dmg-acid2     |   ✅    |
| cpu_instr     |   ✅    |
| mem_timing    |   ✅    |
| mem_timing-2  |   ❌    |
| instr_timing  |   ✅    |


Run script to compile binary

```bash
bash make.sh
```

Execute binary in `build` directory
```bash
./GB-Emulator <path_to_rom>
```